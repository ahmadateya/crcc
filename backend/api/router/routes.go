package router

import (
	"github.com/ahmadateya/crcc/api/handlers/container"

	//"github.com/gin-gonic/gin"
	//"github.com/ahmadateya/crcc/config"
	//"github.com/ahmadateya/crcc/api/router"
	"github.com/ahmadateya/crcc/api/handlers/auth"
	"github.com/gin-gonic/gin"
)

func getRoutes(r *gin.Engine) {
	// dummy routes for test
	getDummyRoutes(r)

	// attaching container routes
	getContainerRoutes(r)
}

func getContainerRoutes(r *gin.Engine) {
	r.GET("/containers", container.List)
}

func getDummyRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.POST("/login", auth.Login)
	r.GET("/me", auth.Me)
}
