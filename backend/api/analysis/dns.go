package analysis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/ahmadateya/crcc/api/models"
)

func DnsScan(dnss string) ([]models.ContainerDns,error){
	var fulldns []string
	var malDnss []models.ContainerDns
	fulldns = strings.Split(dnss, "\n")
	//fulldns=append(fulldns[:len(fulldns)-2],fulldns[len(fulldns)-2+1:]... )
    
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
    isValid:=false
	for _, dns := range fulldns {
		currentDns:= strings.Split(dns, " ")
		checkIp,_:=regexp.Compile("^(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\\.(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])$")
		match := checkIp.FindStringIndex(currentDns[len(currentDns)-1])
        if len(match) == 0 {
			continue
		}

		for _ , malDns := range currentValidDns {
			if currentDns[len(currentDns)-1] != malDns.Dns {
				isValid=true
			}
			
			if (currentDns[len(currentDns)-1] == malDns.Dns) {
				if malDns.White!=true {
				malDnss=append(malDnss,malDns)
				isValid=false
				
				}
				isValid=false
				break
			}
		}
		if isValid {
			malDnss = append(malDnss, models.ContainerDns{Dns: currentDns[len(currentDns)-1],White: false,
				Description: "",Impact: "Medium"})
		}
	}
    //malDnss=append(malDnss[:len(malDnss)-1],malDnss[len(malDnss)-1+1:]... )
	return malDnss, nil
} 