package container

import (
	"encoding/json"
	"fmt"
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
	}
	c.JSON(200, data)
}

func Show(c *gin.Context) {
	containerName := c.Param("container")
	container := containerPkg.GetContainerInfo(containerName)
	var data models.Container
	err := json.Unmarshal([]byte(container), &data)
	fmt.Printf(container)
	if err != nil {
		c.JSON(404, err.Error())
	}
	c.JSON(200, data)
}
