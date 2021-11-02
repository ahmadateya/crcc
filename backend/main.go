package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	//"github.com/ahmadateya/crcc/config"
	"github.com/ahmadateya/crcc/api/router"
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
	router.InitializeConfig(r)
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
	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"token_type":"Bearer",
			"expires_in":31536000,
			"access_token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIyIiwianRpIjoiNWViNTE5OTMxMjgwNjJiZmVjNDk2ZjIyYzNkNmM3ZmIwMDc2OTU1MzczZGEzOWJlOTViOGVlMDVlNTljODA0Njc2MjU4MGY1NGRhMDUxODkiLCJpYXQiOjE2MzU4MDY5NDksIm5iZiI6MTYzNTgwNjk0OSwiZXhwIjoxNjY3MzQyOTQ5LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.lvcyrWJ27e6DA9yLpZHicCphxn2Pcp7himh3CuA5euD5kevaPIW9AegH4WMoPz5E8QfaBfMe-a8akmaO9r6FGvS0iYCEKhKYfamGTrp_NF2t2tyk6LbtTI8r6z0ayukvrx2yKyMS8EdVCc6kCAk0M7aWjiaKI9YCFBaV2DEnCKeyiNcCxAcs_h-cOkxPcEvDfO9J0MY9Dbd57LEgcmKkKXmhIep9M_GTEMvehRH1uGZBVfDyuSGgnImph_NBImc7UhjZTtBbtoSWQ0aSSNjTXzYi5nDfjLzf-c3br698PRMx3BIdlxJt1jM87DkRUCfTf66uNbesBtWuIaFvtFEGpHnVidFwYrUgdmTaAZp4ZlxDvlr-IbDHEKOkbxUTOsLo5LmOcWcli3zhL0Pwtzlf7ce6j7FXeGkUjrNsbTJHWlfnggXEN5-AOTBeeUWgPBToB3BLeFD6vG4d65ZaDMUYbEUYD-tvui64Cfgm0xWjzYlMiqmrx54D4Mib6SXt52fWIP9T1Hou8_EJuXLkr0ICPainlHMqIZpm0VwP0QeCKJNTHJrv2UxKCU_wRQ_ZGwSkcSvb85ZsPoKsQCP4k8rI5_dS9sL1PsRANA7wGJksyjq705PWrOBtMlhvXEp69UL7yEJI8fF1G7S-SJCFYewD0t-RB0o56xWckkjh__8sPMY",
			"refresh_token":"def502008ec8dcdf179a2a9377e012d825dec976361b5a35eb0cb1589b3aab7d78fb1f08956db5004f07133e0483c46a26d1144cfcccec225de204c1f951596ae44a5ee3b7e2fb53331a7cbe69e35c7fbf62f17f89398efe95235d668169486045dee893189ecd606178cbe4b988b57d31c39e997af6d07d7ccdeac09518e649b1cc80b96f4df646d546d0305f66657e10d1bb89fb67a83fc82763f2e42779baa12b6f42ff2242dd63e8335582ec96fd72e6a1627b11895e4426864433064573e7eab44672d11ff5ed9b3ac48f2b281e056421b29916839ddb531052fd809c6c2d8f32724b4bf572e8d647ec520d962b9c669769e2e523ad0c23a18c0222d92e69c431797c8e7e713100092820a80a4f41758074f9bf129f4c364f7e9e37a3bfa85a310d1d68d6ce98175169933e0d77c0acd263d3c1282c62174139b47b25feacf9477e3e01b754b93af6e7588d20d4b33f118968434e22a97b6d788d0d283fed",
		})
	})
	r.GET("/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"type":"users",
			"id":"1",
			"name":"Admin",
			"email":"admin@jsonapi.com",
			"created_at":"2020-12-17 12:32:56",
			"updated_at":"2020-12-17 12:32:56",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
