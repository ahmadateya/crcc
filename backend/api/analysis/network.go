package analysis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ahmadateya/crcc/api/models"
)

func OpenPorts(ports string) ([]models.ContainerPorts,error){
	var fullPorts []string
	var malPorts []models.ContainerPorts
	fullPorts = strings.Split(ports, "\n")

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malPorts.json")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

    var currentMalPorts []map[string]string

	err = json.Unmarshal([]byte(jsonData), &currentMalPorts)

	if err != nil {
		return nil, err
	}

	for _, port := range fullPorts {
		for _, malPort := range currentMalPorts {
            fullPort:= strings.Split(port, " ")[0]
			currentPort := strings.Split(fullPort, ":")
			if   malPort["Port"] == currentPort[len(currentPort)-1] {
				malPorts = append(malPorts, models.ContainerPorts{Port: port,Description: malPort["Description"],
				Impact: malPort["Impact"]})
				break
			}
		}
	}
    
	return malPorts, nil
}