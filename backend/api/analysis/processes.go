package analysis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ahmadateya/crcc/api/models"
)


func ProcessesAnalysisByRegex(processes *models.ContainerProcesses) ([][]string,error){

    var malProcesses [][]string

	getCurrentPath, _:=os.Getwd()
	file, err:= os.Open(getCurrentPath+"/api/analysis/checks/malfilenames.json")

	if err != nil{
		return nil,err
	}
	defer file.Close()
    
	jsonData, _:=ioutil.ReadAll(file)

	var currentMalProcesses map[string]bool
    
	err = json.Unmarshal([]byte(jsonData),&currentMalProcesses)

	if err != nil{
		return nil,err
	}

	for _,process := range processes.Processes{
		for malProcess,mal := range currentMalProcesses{
			re,_ := regexp.Compile(malProcess)
			match := re.FindStringIndex(process[len(process)-1])
			if  mal && len(match) !=0{
				malProcesses = append(malProcesses,process)
			}
		}
	}

	return malProcesses,nil
}