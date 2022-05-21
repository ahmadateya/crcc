package container

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadateya/crcc/api/analysis"
	"github.com/ahmadateya/crcc/api/models"
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
)

func applyFileSystemAnalysis(containerId string) (models.ScanResult, error) {
	// initialize scan result object
	var scanResult models.ScanResult
	scanResult.Title = "File System Analysis"

	// get container files
	containerFiles, err := containerPkg.ListContainerFilesChangesSecondVersion(containerId)
	if err != nil {
		fmt.Printf("Error getting container files: %s\n", err)
		return scanResult, err
	}

	// search for malicious files
	malFiles, err := analysis.FileAnalysisByNameVersionTwo(containerFiles)
	if err != nil {
		fmt.Printf("Error analyzing container files: %s\n", err)
		return scanResult, err
	}

	// formatting scan result
	if len(malFiles) != 0 {
		// just a POC change this later
		var detailsString string
		for _, malFile := range malFiles {
			detailsString += "\n" + malFile.Path
		}
		scanResult.Passed = false
		scanResult.Details = detailsString
	} else {
		scanResult.Passed = true
		scanResult.Details = ""
	}
	return scanResult, nil
}

func applyNetworkAnalysis(containerId string) (models.ScanResult, error) {
	// initialize scan result object
	var scanResult models.ScanResult
	scanResult.Title = "Network Analysis"

	// get container opened ports
	containerPorts, err := containerPkg.ListContainerOpenPorts(containerId)
	fmt.Printf("############## Container ports: %v\n", containerPorts)
	if err != nil {
		fmt.Printf("Error getting container opened ports: %s\n", err)
		return scanResult, err
	}

	// search for malicious ports
	malPorts, err := analysis.CheckOpenPorts(containerPorts)
	if err != nil {
		fmt.Printf("Error analyzing container opened ports: %s\n", err)
		return scanResult, err
	}
	fmt.Printf("===========malPorts======== %+v\n", malPorts)

	// HINT hard coding the malicious ports until scan function is ready
	malPorts = []models.ContainerPorts{
		{"7222", "Lupper worm potentially running on this port", "High"},
		{"4156", "Linux.Slapper.Worm family of worms potentialy running on this port", "High"},
	}
	// formatting scan result
	if len(malPorts) != 0 {
		// just a POC change this later
		var detailsString string
		marshalPorts, _ := json.Marshal(malPorts)
		detailsString = string(marshalPorts)
		scanResult.Passed = false
		scanResult.Details = detailsString
	} else {
		scanResult.Passed = true
		scanResult.Details = ""
	}
	return scanResult, nil
}
