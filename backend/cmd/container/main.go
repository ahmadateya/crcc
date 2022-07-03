package container

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ahmadateya/crcc/api/models"
	"github.com/ahmadateya/crcc/config"
	"github.com/jinzhu/gorm"
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

// ListContainerProcesses Get container running processes
func ListContainerProcesses(containerId string, ps_args string) (models.ContainerProcesses, error) {
	viper := config.NewViper()
	url := fmt.Sprintf("%s:%s/containers/"+containerId+"/top?ps_args="+ps_args, viper.App.Docker.Host, viper.App.Docker.Port)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return models.ContainerProcesses{}, err
	} else {
		p := models.ContainerProcesses{}
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &p)
		return p, nil
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

func ListContainerFilesChangesSecondVersion(containerId string) (string, error) {
	viper := config.NewViper()
	urlForCreatingExec := fmt.Sprintf("%s:%s/containers/%s/exec", viper.App.Docker.Host, viper.App.Docker.Port, containerId)
	createExecRequestData := []byte(`{
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

	startExecRequestData := []byte(`{
  "Detach": false,
  "Tty": false
}`)
	response, err := http.Post(urlForCreatingExec, "application/json", bytes.NewBuffer(createExecRequestData))

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var requestData models.CreateExec
	json.Unmarshal([]byte(body), &requestData)

	urlForStartingExec := fmt.Sprintf("%s:%s/exec/%s/start", viper.App.Docker.Host, viper.App.Docker.Port, requestData.Id)
	response, err = http.Post(urlForStartingExec, "application/json", bytes.NewBuffer(startExecRequestData))

	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ = ioutil.ReadAll(response.Body)

	return string(body), nil
}

func ListContainerOpenPorts(containerId string) (string, error) {

	command := []byte(`{
  "AttachStdin": false,
  "AttachStdout": true,
  "AttachStderr": true,
  "DetachKeys": "ctrl-p,ctrl-q",
  "Tty": false,
  "Cmd": [
    "bash",
      "-c",
    "netstat -anlp | grep -iv 'unix' | awk '{print $4,$7}'"

  ],
  "Env": [
    "FOO=bar",
    "BAZ=quux"
  ]
}`)

	commandOutput, err := RunContainerCommands(containerId, command)
	return commandOutput, err
}

func ListContainerDns(containerId string) (string, error) {
	command := []byte(`{
  "AttachStdin": false,
  "AttachStdout": true,
  "AttachStderr": true,
  "DetachKeys": "ctrl-p,ctrl-q",
  "Tty": false,
  "Cmd": [
    "bash",
	"-c",
    "head -n -1 /etc/resolv.conf"
  ],
  "Env": [
    "FOO=bar",
    "BAZ=quux"
  ]
}`)
	commandOutput, err := RunContainerCommands(containerId, command)
	return commandOutput, err
}

//This function is for running certain commands an a container

func RunContainerCommands(containerId string, command []byte) (string, error) {
	viper := config.NewViper()
	urlForCreatingExec := fmt.Sprintf("%s:%s/containers/%s/exec", viper.App.Docker.Host, viper.App.Docker.Port, containerId)
	createExecRequestData := command

	startExecRequestData := []byte(`{
  "Detach": false,
  "Tty": false
}`)
	response, err := http.Post(urlForCreatingExec, "application/json", bytes.NewBuffer(createExecRequestData))

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	var requestData models.CreateExec
	json.Unmarshal([]byte(body), &requestData)
    
	urlForStartingExec := fmt.Sprintf("%s:%s/exec/%s/start", viper.App.Docker.Host, viper.App.Docker.Port, requestData.Id)
	response, err = http.Post(urlForStartingExec, "application/json", bytes.NewBuffer(startExecRequestData))

	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, _ = ioutil.ReadAll(response.Body)
    
	return string(body), nil
}

func ListContainerHistory(containerId string) []config.Container {
	// connect to the database
	viper := config.NewViper()
	db, err := gorm.Open("postgres", viper.Database.Connection)
	if err != nil {
		fmt.Printf(err.Error())
		panic(err.Error())
	}
	defer db.Close()
	var containerScans []config.Container // DB model
	db.Where("name = ?", containerId).Find(&containerScans)
	return containerScans
}
