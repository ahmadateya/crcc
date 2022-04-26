package analysis

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ahmadateya/crcc/api/models"
)

func FileAnalysisByName(files *[]models.FileInfo) ([]models.FileInfo,error){
    
	
	var malFiles []models.FileInfo
	
    getCurrentPath, _:=os.Getwd()
	file, err:= os.Open(getCurrentPath+"/api/analysis/checks/malfilenames.json")

	if err != nil{
		return nil,err
	}
	defer file.Close()
    
	jsonData, _:=ioutil.ReadAll(file)

	var currentMalFiles map[string]bool
    
	err = json.Unmarshal([]byte(jsonData),&currentMalFiles)
    
	if err != nil{
		return nil,err
	}

	for _,file := range *files{
		for malFile,mal := range currentMalFiles{
			path:=strings.Split(file.Path, "/")
			if malFile == path[len(path)-1] && mal{
				malFiles = append(malFiles, models.FileInfo{Path: file.Path,Kind: file.Kind})
			}
		}
	}
	
	return malFiles,nil

}

func FileAnalysisByNameVersionTwo(files string)([]models.FileInfo,error){
	var fullPath []string
	var malFiles []models.FileInfo

	fullPath=strings.Split(files,"\n")
    
	
    
    

    getCurrentPath, _:=os.Getwd()
	file, err:= os.Open(getCurrentPath+"/api/analysis/checks/malfilenames.json")

	if err != nil{
		return nil,err
	}
	defer file.Close()
    
	jsonData, _:=ioutil.ReadAll(file)

	var currentMalFiles map[string]bool
    
	err = json.Unmarshal([]byte(jsonData),&currentMalFiles)
    
	if err != nil{
		return nil,err
	}

	for _,file:= range fullPath{
		for malFile,mal := range currentMalFiles{
			path:=strings.Split(file, "/")
			if mal && malFile == path[len(path)-1]{
				malFiles = append(malFiles, models.FileInfo{Path: file,Kind: 0})
			}
		}
	}
	
	return malFiles,nil
}