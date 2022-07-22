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
		marshalPorts, _ := json.Marshal(malFiles)
		detailsString = string(marshalPorts)
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
    fmt.Println(containerPorts)
	// search for malicious ports
	malPorts, err := analysis.CheckOpenPorts(containerPorts)
	if err != nil {
		fmt.Printf("Error analyzing container opened ports: %s\n", err)
		return scanResult, err
	}
	fmt.Printf("===========malPorts======== %+v\n", malPorts)

	// HINT hard coding the malicious ports until scan function is ready
	//malPorts = []models.ContainerPorts{
	//	{"7222", "Lupper worm potentially running on this port", "High"},
	//	{"4156", "Linux.Slapper.Worm family of worms potentialy running on this port", "High"},
	//}
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

func applyProcessAnalysis(containerId string) (models.ScanResult, error) {
	// initialize scan result object
	var scanResult models.ScanResult
	scanResult.Title = "Process Analysis"

	// get container processes
	containerProcesses, err := containerPkg.ListContainerProcesses(containerId, "")
	if err != nil {
		fmt.Printf("Error getting container processes: %s\n", err)
		return scanResult, err
	}

	fmt.Printf("############## Container processes: %v\n", containerProcesses)
	// search for malicious processes
	malProcesses, err := analysis.ProcessesAnalysisByRegex(containerProcesses)
	if err != nil {
		fmt.Printf("Error analyzing container processes: %s\n", err)
		return scanResult, err
	}

	// formatting scan result
	if len(malProcesses) != 0 {
		// just a POC change this later
		var detailsString string
		marshalPorts, _ := json.Marshal(malProcesses)
		detailsString = string(marshalPorts)
		scanResult.Passed = false
		scanResult.Details = detailsString
	} else {
		scanResult.Passed = true
		scanResult.Details = ""
	}
	return scanResult, nil
}

func applyDnsAnalysis(containerId string) (models.ScanResult, error) {
	// initialize scan result object
	var scanResult models.ScanResult
	scanResult.Title = "DNS Analysis"

	// get container processes
	contaienrDns, err := containerPkg.ListContainerDns(containerId)
	if err != nil {
		fmt.Printf("Error getting container dns: %s\n", err)
		return scanResult, err
	}

	fmt.Printf("############## Container dns: %v\n", contaienrDns)
	// search for malicious processes
	malDns, err := analysis.DnsScan(contaienrDns)
	if err != nil {
		fmt.Printf("Error analyzing container dns: %s\n", err)
		return scanResult, err
	}

	// formatting scan result
	if len(malDns) != 0 {
		// just a POC change this later
		var detailsString string
		marshalDns, _ := json.Marshal(malDns)
		detailsString = string(marshalDns)
		scanResult.Passed = false
		scanResult.Details = detailsString
	} else {
		scanResult.Passed = true
		scanResult.Details = ""
	}
	return scanResult, nil
}

//func applyDNSAnalysis(containerId string) (models.ScanResult, error) {
//}
