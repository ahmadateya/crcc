package container

import (
	"fmt"
	"github.com/ahmadateya/crcc/config"
	"io/ioutil"
	"net/http"
)

// Container package to handle container related business

// ListContainers list all containers
func ListContainers() string {
	viper := config.NewViper()
	url := fmt.Sprintf("%s:%s/containers/json", viper.App.Docker.Host, viper.App.Docker.Port)
	fmt.Printf("================= url %+v \n", url)

	response, err := http.Get(url)
	fmt.Printf("================= %+v\n", response)
	fmt.Printf("================= err %+v\n", err)

	if err != nil {
		return fmt.Sprintf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
}

// GetContainerInfo get container info
func GetContainerInfo(containerId string) string {
	viper := config.NewViper()
	url := fmt.Sprintf("%s:%s/containers/%s/json", viper.App.Docker.Host, viper.App.Docker.Port, containerId)
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
}
