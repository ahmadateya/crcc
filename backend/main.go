package main

import (
	"fmt"
	"github.com/ahmadateya/crcc/api/router"
	"github.com/ahmadateya/crcc/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// initiating the config file
	viper := config.NewViper()

	// initialize database object
	db, err := gorm.Open("postgres", viper.Database.Connection)
	if err != nil {
		fmt.Printf(err.Error())
		panic(err.Error())
	}

	// migrate the schema
	m := db.AutoMigrate(config.Container{})

	if m != nil && m.Error != nil {
		//We have an error
		fmt.Printf(m.Error.Error())
	}
	// close the database connection after the migration
	db.Close()

	// initiating the routes and router config
	r := router.Init()
	err = r.Run(viper.Server.Port)
	if err != nil {
		return
	}
}
