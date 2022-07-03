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
	//getDummyRoutes(r)

	// attaching container routes
	getContainerRoutes(r)
}

func getContainerRoutes(r *gin.Engine) {
	r.GET("/containers", container.List)
	r.GET("/containers/:container", container.Show)
	r.GET("/containers/:container/scan", container.Scan)
	r.GET("/containers/:container/history", container.History)
	r.GET("/processesrules",container.GetProcesses)
	r.GET("/filesrules",container.GetFiles)
	r.GET("/networkrules",container.GetPorts)
	r.GET("/dnsrules",container.GetDNS)
	r.POST("/addprocess",container.AddProcess)
	r.POST("/addfile",container.AddFile)
	r.POST("/addport",container.AddPort)
	r.POST("/adddns",container.AddDns)
	r.DELETE("/deleteprocess/:index",container.DeleteProcess)
	r.DELETE("/deletefile/:index",container.DeleteFile)
    r.DELETE("/deleteport/:index",container.DeletePort)
	r.DELETE("/deleteDns/:index",container.DeleteDns)
	r.PUT("/editprocess/:index",container.EditProcess)
	r.PUT("/editfile/:index",container.EditFile)
	r.PUT("/editport/:index",container.EditPort)
	r.PUT("/editdns/:index",container.EditDns)
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
