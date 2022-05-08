package container

import (
	"fmt"
	"github.com/ahmadateya/crcc/api/analysis"
	"github.com/ahmadateya/crcc/api/models"
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
)

func applyFileSystemAnalysis(containerId string) (models.ScanResult, error) {
	// get container files
	containerFiles, err := containerPkg.ListContainerFilesChangesSecondVersion(containerId)
	var scanResult models.ScanResult
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

	if len(malFiles) != 0 {
		// dummy crappy code to create the details, just POC change this later
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
