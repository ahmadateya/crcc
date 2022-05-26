package container

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadateya/crcc/api/analysis"
	"github.com/ahmadateya/crcc/api/models"
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
	"github.com/ahmadateya/crcc/config"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"math"
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

	// calc the compliance
	scanResponse.Compliance = calcCompliance(scanResponse.Results)

	// store the analysis in the database
	err = storeAnalysis(containerId, scanResponse)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
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

	containerRecord := config.Container{
		Name: containerId,
		Scan: string(scanJsonData),
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
