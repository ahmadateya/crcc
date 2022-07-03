package analysis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ahmadateya/crcc/api/models"
)

func DnsScan(dnss string) ([]models.ContainerDns,error){
	var fulldns []string
	var malDnss []models.ContainerDns
	fulldns = strings.Split(dnss, "\n")
    
	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malDns.json")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

    var currentValidDns []models.ContainerDns

	err = json.Unmarshal([]byte(jsonData), &currentValidDns)

	if err != nil {
		return nil, err
	}

	for _, dns := range fulldns {
		currentDns:= strings.Split(dns, " ")
		
		for _ , malDns := range currentValidDns {
			if currentDns[len(currentDns)-1] != malDns.Dns {
				malDnss = append(malDnss, models.ContainerDns{Dns: currentDns[len(currentDns)-1],White: false,
				Description: "",Impact: "Medium"})
			}else if (currentDns[len(currentDns)-1] == malDns.Dns && malDns.White!=true) {
				malDnss=append(malDnss,malDns)
			}
		}
		
	}
    
	return malDnss, nil
} 