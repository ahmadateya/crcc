package container

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ahmadateya/crcc/api/analysis"
	"github.com/ahmadateya/crcc/api/models"
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
	"github.com/ahmadateya/crcc/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func List(c *gin.Context) {
	containers := containerPkg.ListContainers()
	// TODO : handle the integers in the JSON response, it doesnt work
	var data []models.ListContainer
	err := json.Unmarshal([]byte(containers), &data)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, data)
}

func Show(c *gin.Context) {
	containerName := c.Param("container")
	container := containerPkg.GetContainerInfo(containerName)
	var data models.Container
	err := json.Unmarshal([]byte(container), &data)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, data)
}

func ListProcesses(c *gin.Context) {
	containerId := c.Param("container")
	processes, err := containerPkg.ListContainerProcesses(containerId, "")
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	c.JSON(200, processes)
}

func ListFileChanges(c *gin.Context) {
	containerName := c.Param("container")
	container := containerPkg.ListContainerFilesChanges(containerName)
	var data []models.FileInfo
	err := json.Unmarshal([]byte(container), &data)
	if err != nil {
		c.JSON(404, err.Error())
	}

	malFiles, err := analysis.FileAnalysisByName(&data)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, malFiles)
}

func ListFileChangesWithNameVersionTwo(c *gin.Context) {
	containerId := c.Param("container")
	container, err := containerPkg.ListContainerFilesChangesSecondVersion(containerId)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	malFiles, err := analysis.FileAnalysisByNameVersionTwo(container)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, malFiles)
}

func ListNetworks(c *gin.Context) {
	containerId := c.Param("container")
	container := containerPkg.ListContainerNetworks(containerId)
	var data map[string]interface{}
	err := json.Unmarshal([]byte(container), &data)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	containerPkg.ListContainerFilesChangesSecondVersion(containerId)
	c.JSON(200, data["NetworkSettings"])
}

func Scan(c *gin.Context) {
	containerId := c.Param("container")
	var scanResponse models.ScanDataResponse

	// apply filesystem analysis to the container
	fileSystemScan, err := applyFileSystemAnalysis(containerId)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	// append file system analysis to the scan response
	scanResponse.Results = append(scanResponse.Results, fileSystemScan)

	// apply network analysis to the container
	networkScan, err := applyNetworkAnalysis(containerId)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	// append network analysis to the scan response
	scanResponse.Results = append(scanResponse.Results, networkScan)

	// apply process analysis to the container
	processScan, err := applyProcessAnalysis(containerId)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	// append process analysis to the scan response
	scanResponse.Results = append(scanResponse.Results, processScan)

	// apply dns analysis to the container
	//dnsScan, err := applyDnsAnalysis(containerId)
	//if err != nil {
	//	c.JSON(500, err.Error())
	//	return
	//}
	// append dns analysis to the scan response
	//scanResponse.Results = append(scanResponse.Results, dnsScan)
	// calc the compliance
	scanResponse.Compliance = calcCompliance(scanResponse.Results)

	// store the analysis in the database
	/*err = storeAnalysis(containerId, scanResponse)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}*/

	// return the response
	c.JSON(200, scanResponse)
}

// very basic equation to calculate the compliance
func calcCompliance(results []models.ScanResult) float32 {
	var passed, failed int
	var compliance float32
	for _, result := range results {
		if result.Passed {
			passed = passed + 1
		} else {
			failed = failed + 1
		}
	}
	compliance = (float32(passed) * float32(100)) / float32(len(results))
	return float32(math.Round(float64(compliance*100)) / 100)
}

// store the results of the analysis in the database
func storeAnalysis(containerId string, results models.ScanDataResponse) error {
	// connect to the database
	viper := config.NewViper()
	db, err := gorm.Open("postgres", viper.Database.Connection)
	if err != nil {
		fmt.Printf(err.Error())
		panic(err.Error())
	}
	defer db.Close()

	scanJsonData, err := json.Marshal(results)
	if err != nil {
		return err
	}

	loc, _ := time.LoadLocation("Africa/Cairo")
	containerRecord := config.Container{
		Name:      containerId,
		Scan:      string(scanJsonData),
		CreatedAt: time.Now().In(loc).Format("2006-01-02 3:4:5 pm"),
	}
	db.Create(&containerRecord) // pass pointer of data to Create
	return nil
}

// History returns the history of the container scans
func History(c *gin.Context) {
	containerId := c.Param("container")
	containerScans := containerPkg.ListContainerHistory(containerId)
	c.JSON(200, containerScans)
}

func StoreCNN(c *gin.Context) {
	containerId := c.Query("container")

	type Data struct {
		Data models.ScanDataResponse `json:"data"`
	}
	var data Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := storeAnalysis(containerId, data.Data)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, `{}`)

}

func GetProcesses(c *gin.Context) {
	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malProcesses.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalProcesses []models.ContainerProcess

	err = json.Unmarshal([]byte(jsonData), &currentMalProcesses)

	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, currentMalProcesses)
}

func GetFiles(c *gin.Context) {
	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malfilenames.json")

	if err != nil {
		c.JSON(400, err)

		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalFiles []models.ContainerFile

	err = json.Unmarshal([]byte(jsonData), &currentMalFiles)

	if err != nil {
		c.JSON(400, err)

		return
	}
	c.JSON(200, currentMalFiles)
}

func GetDNS(c *gin.Context) {
	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malDns.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalDns []models.ContainerDns

	err = json.Unmarshal([]byte(jsonData), &currentMalDns)

	if err != nil {
		c.JSON(400, err)
		return
	}
	c.JSON(200, currentMalDns)
}

func GetPorts(c *gin.Context) {
	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malPorts.json")

	if err != nil {
		c.JSON(400, err)

		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalPorts []models.ContainerPorts

	err = json.Unmarshal([]byte(jsonData), &currentMalPorts)

	if err != nil {
		c.JSON(400, err)

		return
	}
	c.JSON(200, currentMalPorts)
}

func AddProcess(c *gin.Context) {

	var data models.ContainerProcess
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if data.Cmd == "" || data.Impact == "" {
		c.JSON(400, `{"err": "enter cmd or impact"}`)
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malProcesses.json")

	if err != nil {
		c.JSON(400, err)

		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalProcesses []models.ContainerProcess

	err = json.Unmarshal([]byte(jsonData), &currentMalProcesses)

	if err != nil {
		c.JSON(400, err)

		return
	}
	currentMalProcesses = append(currentMalProcesses, data)

	jsonData, err = json.Marshal(currentMalProcesses)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malProcesses.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalProcesses)
}

func AddFile(c *gin.Context) {

	var data models.ContainerFile
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if data.File == "" || data.Impact == "" {
		c.JSON(400, `{"err": "enter File or Impact"}`)
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malfilenames.json")

	if err != nil {
		c.JSON(400, err)

		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalFiles []models.ContainerFile

	err = json.Unmarshal([]byte(jsonData), &currentMalFiles)

	if err != nil {
		c.JSON(400, err)

		return
	}
	currentMalFiles = append(currentMalFiles, data)

	jsonData, err = json.Marshal(currentMalFiles)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malfilenames.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalFiles)
}

func AddDns(c *gin.Context) {

	var data models.ContainerDns
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if data.Dns == "" {
		c.JSON(400, `{"err": "enter Dns or impact"}`)
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malDns.json")

	if err != nil {
		c.JSON(400, err)

		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalDns []models.ContainerDns

	err = json.Unmarshal([]byte(jsonData), &currentMalDns)

	if err != nil {
		c.JSON(400, err)

		return
	}
	currentMalDns = append(currentMalDns, data)

	jsonData, err = json.Marshal(currentMalDns)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malDns.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalDns)
}

func AddPort(c *gin.Context) {

	var data models.ContainerPorts
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if data.Port == "" || data.Impact == "" {
		c.JSON(400, `{"err": "enter port or impact"}`)
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malPorts.json")

	if err != nil {
		c.JSON(400, err)

		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalPorts []models.ContainerPorts

	err = json.Unmarshal([]byte(jsonData), &currentMalPorts)

	if err != nil {
		c.JSON(400, err)

		return
	}
	currentMalPorts = append(currentMalPorts, data)

	jsonData, err = json.Marshal(currentMalPorts)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malPorts.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalPorts)
}

func DeleteProcess(c *gin.Context) {
	index := c.Param("index")

	if index == "" {
		c.JSON(400, "index is not integer")
		return
	}

	dataIndex, _ := strconv.Atoi(index)

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malProcesses.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalProcesses []models.ContainerProcess

	err = json.Unmarshal([]byte(jsonData), &currentMalProcesses)

	if err != nil {
		c.JSON(400, err)
		return
	}
	currentMalProcesses = append(currentMalProcesses[:dataIndex], currentMalProcesses[dataIndex+1:]...)

	jsonData, err = json.Marshal(currentMalProcesses)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malProcesses.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalProcesses)
}

func DeleteFile(c *gin.Context) {
	index := c.Param("index")

	if index == "" {
		c.JSON(400, "index is not integer")
		return
	}

	dataIndex, _ := strconv.Atoi(index)

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malfilenames.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalFiles []models.ContainerFile

	err = json.Unmarshal([]byte(jsonData), &currentMalFiles)

	if err != nil {
		c.JSON(400, err)
		return
	}
	currentMalFiles = append(currentMalFiles[:dataIndex], currentMalFiles[dataIndex+1:]...)

	jsonData, err = json.Marshal(currentMalFiles)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malfilenames.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalFiles)
}

func DeletePort(c *gin.Context) {
	index := c.Param("index")

	if index == "" {
		c.JSON(400, "index is not integer")
		return
	}

	dataIndex, _ := strconv.Atoi(index)

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malPorts.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalPorts []models.ContainerPorts

	err = json.Unmarshal([]byte(jsonData), &currentMalPorts)

	if err != nil {
		c.JSON(400, err)
		return
	}
	currentMalPorts = append(currentMalPorts[:dataIndex], currentMalPorts[dataIndex+1:]...)

	jsonData, err = json.Marshal(currentMalPorts)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malPorts.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalPorts)
}

func DeleteDns(c *gin.Context) {
	index := c.Param("index")

	if index == "" {
		c.JSON(400, "index is not integer")
		return
	}

	dataIndex, _ := strconv.Atoi(index)

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malDns.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalDns []models.ContainerDns

	err = json.Unmarshal([]byte(jsonData), &currentMalDns)

	if err != nil {
		c.JSON(400, err)
		return
	}
	currentMalDns = append(currentMalDns[:dataIndex], currentMalDns[dataIndex+1:]...)

	jsonData, err = json.Marshal(currentMalDns)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malDns.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalDns)
}

func EditFile(c *gin.Context) {
	index := c.Param("index")

	currentIndex, _ := strconv.Atoi(index)

	var data models.ContainerFile
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malfilenames.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalFiles []models.ContainerFile

	err = json.Unmarshal([]byte(jsonData), &currentMalFiles)

	if err != nil {
		c.JSON(400, err)
		return
	}
	if currentIndex < 0 || currentIndex >= len(currentMalFiles) {
		c.JSON(400, "Rule doesn't exist")
	}
	currentMalFiles[currentIndex] = data
	jsonData, err = json.Marshal(currentMalFiles)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malfilenames.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalFiles)

}

func EditPort(c *gin.Context) {
	index := c.Param("index")

	currentIndex, _ := strconv.Atoi(index)

	var data models.ContainerPorts
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malPorts.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalPorts []models.ContainerPorts

	err = json.Unmarshal([]byte(jsonData), &currentMalPorts)

	if err != nil {
		c.JSON(400, err)
		return
	}
	if currentIndex < 0 || currentIndex >= len(currentMalPorts) {
		c.JSON(400, "Rule doesn't exist")
	}
	currentMalPorts[currentIndex] = data
	jsonData, err = json.Marshal(currentMalPorts)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malPorts.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalPorts)

}

func EditProcess(c *gin.Context) {
	index := c.Param("index")

	currentIndex, _ := strconv.Atoi(index)

	var data models.ContainerProcess
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malProcesses.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalProcesses []models.ContainerProcess

	err = json.Unmarshal([]byte(jsonData), &currentMalProcesses)

	if err != nil {
		c.JSON(400, err)
		return
	}
	if currentIndex < 0 || currentIndex >= len(currentMalProcesses) {
		c.JSON(400, "Rule doesn't exist")
	}
	currentMalProcesses[currentIndex] = data
	jsonData, err = json.Marshal(currentMalProcesses)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malProcesses.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalProcesses)

}

func EditDns(c *gin.Context) {
	index := c.Param("index")

	currentIndex, _ := strconv.Atoi(index)

	var data models.ContainerDns
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malDns.json")

	if err != nil {
		c.JSON(400, err)
		return
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalDns []models.ContainerDns

	err = json.Unmarshal([]byte(jsonData), &currentMalDns)

	if err != nil {
		c.JSON(400, err)
		return
	}
	if currentIndex < 0 || currentIndex >= len(currentMalDns) {
		c.JSON(400, "Rule doesn't exist")
	}
	currentMalDns[currentIndex] = data
	jsonData, err = json.Marshal(currentMalDns)

	err = ioutil.WriteFile(getCurrentPath+"/api/analysis/checks/malDns.json", jsonData, 0644)

	if err != nil {
		c.JSON(400, err)

		return
	}

	c.JSON(200, currentMalDns)

}
