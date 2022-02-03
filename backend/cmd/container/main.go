package container

import (
	"fmt"
	"github.com/ahmadateya/crcc/config"
	"io/ioutil"
	"net/http"
)

// container package to handle container related business

// TODO: make a connection to docker daemon on the host machine

func ListContainers() string {
	viper := config.NewViper()
	url := fmt.Sprintf("%s:%s/containers/json", viper.App.Docker.Host, viper.App.Docker.Port)
	fmt.Println("Starting the application...")
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return ""
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		return string(data)
	}
	//jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	//jsonValue, _ := json.Marshal(jsonData)
}
