package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/ahmadateya/crcc/config"
	//"github.com/ahmadateya/crcc/api/router"
)

func main() {
	//viper := config.NewViper()
	//var newDB db.Database
	//newDB = db.NewPostgres()
	//obj := newDB.Open()
	//defer obj.Close()
	//
	//router := router.Init()
	//InitializeConfig(router)
	//
	//v1 := router.Group("/api/v1")
	//InitializeRouts(obj, v1)
	//err = router.Run(viper.Server.Port)
	//if err != nil {
	//	fmt.Println(err)
	//}


	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
