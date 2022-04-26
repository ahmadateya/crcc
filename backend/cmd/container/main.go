package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ahmadateya/crcc/api/models"
	"github.com/ahmadateya/crcc/config"
)

// Container package to handle container related business

// ListContainers list all containers
func ListContainers() string {
	viper := config.NewViper()
	url := fmt.Sprintf("%s:%s/containers/json", viper.App.Docker.Host, viper.App.Docker.Port)
	response, err := http.Get(url)
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

// Get container running processes
func ListContainerProcesses(containerId string, ps_args string) string {
	viper := config.NewViper()
	url := fmt.Sprintf("%s:%s/containers/"+containerId+"/top?ps_args="+ps_args, viper.App.Docker.Host, viper.App.Docker.Port)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
}


func ListContainerFilesChanges(containerId string) string {
	viper := config.NewViper()
	url := fmt.Sprintf("%s:%s/containers/%s/changes", viper.App.Docker.Host, viper.App.Docker.Port, containerId)
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		return string(data)
	}
}

func ListContainerNetworks(containerId string) string {
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

func ListContainerFilesChangesSecondVersion(containerId string) (string,error){
	viper := config.NewViper()
	urlForCreatingExec := fmt.Sprintf("%s:%s/containers/%s/exec", viper.App.Docker.Host, viper.App.Docker.Port, containerId)
	createExecRequestData:=[]byte(`{
  "AttachStdin": false,
  "AttachStdout": true,
  "AttachStderr": true,
  "DetachKeys": "ctrl-p,ctrl-q",
  "Tty": false,
  "Cmd": [
    "find",
	"/"
  ],
  "Env": [
    "FOO=bar",
    "BAZ=quux"
  ]
}`)

    startExecRequestData:=[]byte(`{
  "Detach": false,
  "Tty": false
}`)
	response, err := http.Post(urlForCreatingExec,"application/json", bytes.NewBuffer(createExecRequestData))
	
	if err != nil {
		return "",err
	}
	defer response.Body.Close()
    
	body, _ := ioutil.ReadAll(response.Body)
	var requestData models.CreateExec
	json.Unmarshal([]byte(body),&requestData)

	urlForStartingExec := fmt.Sprintf("%s:%s/exec/%s/start", viper.App.Docker.Host, viper.App.Docker.Port,requestData.Id)
	response, err=http.Post(urlForStartingExec,"application/json",bytes.NewBuffer(startExecRequestData))

	if err!=nil{
		return "",err
	}
    defer response.Body.Close()
	body,_=ioutil.ReadAll(response.Body)
    
	return string(body),nil
}

