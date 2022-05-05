package analysis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/ahmadateya/crcc/api/models"
)


type MalProcesses struct{
	Cmd string `json:"cmd"`
	User string `json:"user"`
	Type byte `json:"type"`
	Applied bool `json:"applied"`
}

func ProcessesAnalysisByRegex(processes *models.ContainerProcesses) ([][]string,error){

    var malProcesses [][]string

	getCurrentPath, _:=os.Getwd()
	file, err:= os.Open(getCurrentPath+"/api/analysis/checks/malfilenames.json")

	if err != nil{
		return nil,err
	}
	defer file.Close()
    
	jsonData, _:=ioutil.ReadAll(file)

	var currentMalProcesses []MalProcesses
    
	err = json.Unmarshal([]byte(jsonData),&currentMalProcesses)

	if err != nil{
		return nil,err
	}

	for _,process := range processes.Processes{
		for _,malProcess := range currentMalProcesses{
			re,_ := regexp.Compile(malProcess.Cmd)
			match := re.FindStringIndex(process[len(process)-1])
			if  malProcess.Applied && len(match) !=0{
				checkProcessTypeScan(&malProcess,&malProcesses,&process)
			}
		}
	}



	return malProcesses,nil
}

func checkProcessTypeScan(malProcess *MalProcesses,malProcesses *[][]string,process *[]string){
	if malProcess.Type == 0{
					(*malProcesses) = append((*malProcesses),*process)
				}else if malProcess.Type == 1{
					if malProcess.User != (*process)[0]{
						(*malProcesses) = append((*malProcesses),*process)
					}
				}
}