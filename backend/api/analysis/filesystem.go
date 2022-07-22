package analysis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/ahmadateya/crcc/api/models"
)

func FileAnalysisByName(files *[]models.FileInfo) ([]models.ContainerFile, error) {

	var malFiles []models.ContainerFile

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malfilenames.json")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalFiles []models.ContainerFile

	err = json.Unmarshal([]byte(jsonData), &currentMalFiles)

	if err != nil {
		return nil, err
	}

	for _, file := range *files {
		path := strings.Split(file.Path, "/")
		for _, malFile := range currentMalFiles{
			if malFile.File == path[len(path)-1] {
				malFiles = append(malFiles, models.ContainerFile{File: path[len(path)-1],
				Description: malFile.Description, Impact: malFile.Impact})
				break
			}
		}
	}

	return malFiles, nil

}

func FileAnalysisByNameVersionTwo(files string) ([]models.ContainerFile, error) {
	var fullPath []string
	var malFiles []models.ContainerFile

	fullPath = strings.Split(files, "\n")

	getCurrentPath, _ := os.Getwd()
	file, err := os.Open(getCurrentPath + "/api/analysis/checks/malfilenames.json")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, _ := ioutil.ReadAll(file)

	var currentMalFiles []models.ContainerFile

	err = json.Unmarshal([]byte(jsonData), &currentMalFiles)

	if err != nil {
		return nil, err
	}

	for _, file := range fullPath {
		path := strings.Split(file, "/")
		for _, malFile := range currentMalFiles{
			re, _ := regexp.Compile(malFile.File)
			match := re.FindStringIndex(path[len(path)-1])
			if len(match) != 0 {
				malFiles = append(malFiles, models.ContainerFile{File: path[len(path)-1],
				Description: malFile.Description, Impact: malFile.Impact})
				break
			}
		}
	}

	fmt.Printf("=================================== %+v\n", malFiles)
	return malFiles, nil
}
