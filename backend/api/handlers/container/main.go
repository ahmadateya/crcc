package container

import (
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	containers := containerPkg.ListContainers()
	c.JSON(200, gin.H{
		"containers ": containers,
	})
}
