package analysis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ahmadateya/crcc/api/models"
)

//The process.type here is for checking if the user wants to check the process.cmd it self or he just wants to see
// if the current command is running under a certain user


func ProcessesAnalysisByRegex(processes *models.ContainerProcesses) ([]models.ContainerProcess,error){

    var malProcesses []models.ContainerProcess

	getCurrentPath, _:=os.Getwd()
	file, err:= os.Open(getCurrentPath+"/api/analysis/checks/malProcesses.json")

	if err != nil{
		return nil,err
	}
	defer file.Close()
    
	jsonData, _:=ioutil.ReadAll(file)

	var currentMalProcesses []models.ContainerProcess
    
	err = json.Unmarshal([]byte(jsonData),&currentMalProcesses)

	if err != nil{
		return nil,err
	}

	for _,process := range processes.Processes{
		for _,malProcess := range currentMalProcesses{
			re,_ := regexp.Compile(malProcess.Cmd)
			match := re.FindStringIndex(process[len(process)-1])
			if  len(match)!=0{
				checkProcessTypeScan(&malProcess,&malProcesses,&process)
			}
		}
	}



	return malProcesses,nil
}

func checkProcessTypeScan(malProcess *models.ContainerProcess,malProcesses *[]models.ContainerProcess,process *[]string){
	if malProcess.Type == 0{
					(*malProcesses) = append((*malProcesses),models.ContainerProcess{
						Cmd: (*process)[7],
						User: (*process)[0],
						Type: malProcess.Type,
						Description: malProcess.Description,
						Impact: malProcess.Impact,
					})
				}else if malProcess.Type == 1{
					if malProcess.User != (*process)[0]{
						(*malProcesses) = append((*malProcesses),models.ContainerProcess{
						Cmd: (*process)[7],
						User: (*process)[0],
						Type: malProcess.Type,
						Description: malProcess.Description,
						Impact: malProcess.Impact,
					})
					}
				}
}