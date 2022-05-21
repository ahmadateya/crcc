package container

import (
	"fmt"
	"github.com/ahmadateya/crcc/api/analysis"
	"github.com/ahmadateya/crcc/api/models"
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
)

func applyFileSystemAnalysis(containerId string) (models.ScanResult, error) {
	// initialize scan result object
	var scanResult models.ScanResult

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

		scanResult.Title = "File System Analysis"
		scanResult.Passed = false
		scanResult.Details = detailsString
	} else {
		scanResult.Title = "File System Analysis"
		scanResult.Passed = true
		scanResult.Details = ""
	}
	return scanResult, nil
}

func applyNetworkAnalysis(containerId string) (models.ScanResult, error) {
	// initialize scan result object
	var scanResult models.ScanResult
	// get container opened ports
	containerPorts, err := containerPkg.ListContainerOpenPorts(containerId)
	if err != nil {
		fmt.Printf("Error getting container opened ports: %s\n", err)
		return scanResult, err
	}

	fmt.Printf("================== %+v\n", containerPorts)
	// search for malicious ports
	malPorts, err := analysis.CheckOpenPorts(containerPorts)
	if err != nil {
		fmt.Printf("Error analyzing container opened ports: %s\n", err)
		return scanResult, err
	}

	// formatting scan result
	if len(malPorts) != 0 {
		// just a POC change this later
		var detailsString string
		for _, malPort := range malPorts {
			detailsString += "\n" + fmt.Sprintf("%#v", malPort)
		}
		scanResult.Title = "Network Analysis"
		scanResult.Passed = false
		scanResult.Details = detailsString
	} else {
		scanResult.Title = "Network Analysis"
		scanResult.Passed = true
		scanResult.Details = ""
	}
	return scanResult, nil
}
