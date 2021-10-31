package config

import (
	"github.com/spf13/viper"
	"log"
)

func NewViper() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/var/www/go")
	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return &configuration
}
