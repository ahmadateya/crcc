package main

import (

	//"github.com/ahmadateya/crcc/config"
	"github.com/ahmadateya/crcc/api/router"
)

func main() {
	//viper := config.NewViper()
	//var newDB db.Database
	//newDB = db.NewPostgres()
	//obj := newDB.Open()
	//defer obj.Close()
	//InitializeConfig(router)
	//
	//v1 := router.Group("/api/v1")
	//InitializeRouts(obj, v1)
	//err = router.Run(viper.Server.Port)
	//if err != nil {
	//	fmt.Println(err)
	//}

	r := router.Init()
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		return
	}
}
