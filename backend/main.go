package main

import (

	//"github.com/ahmadateya/crcc/config"
	"github.com/ahmadateya/crcc/api/router"
	"github.com/ahmadateya/crcc/config"
)

func main() {
	//var newDB db.Database
	//newDB = db.NewPostgres()
	//obj := newDB.Open()
	//defer obj.Close()
	//InitializeConfig(router)
	//v1 := router.Group("/api/v1")
	//InitializeRouts(obj, v1)
	//err = router.Run(viper.Server.Port)
	//if err != nil {
	//	fmt.Println(err)
	//}

	// initiating the config file
	viper := config.NewViper()

	// initiating the routes and router config
	r := router.Init()
	err := r.Run(viper.Server.Port)
	if err != nil {
		return
	}
}
