package container

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadateya/crcc/api/analysis"
	"github.com/ahmadateya/crcc/api/models"
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
	"github.com/gin-gonic/gin"
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
	container := containerPkg.ListContainerProcesses(containerId,"")
	fmt.Println(container)
	var data models.ContainerProcesses
	err := json.Unmarshal([]byte(container), &data)
	if err != nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, data)
}

func ListFileChanges(c *gin.Context) {
	containerName := c.Param("container")
	container := containerPkg.ListContainerFilesChanges(containerName)
	var data []models.FileInfo
	err := json.Unmarshal([]byte(container), &data)
	if err != nil {
		c.JSON(404, err.Error())
	}

	malFiles, err:=analysis.FileAnalysisByName(&data)
	if err !=nil {
		c.JSON(404, err.Error())
		return
	}
	c.JSON(200, malFiles)
}

func ListFileChangesWithNameVersionTwo(c *gin.Context) {
	containerId := c.Param("container")
	container,err := containerPkg.ListContainerFilesChangesSecondVersion(containerId)
	if err !=nil {
		c.JSON(404, err.Error())
		return
	}
	malFiles, err:=analysis.FileAnalysisByNameVersionTwo(container)
	if err !=nil {
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
	// TODO:: trigger the scan
}
