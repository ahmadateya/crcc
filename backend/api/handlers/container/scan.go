package container

import (
	"github.com/ahmadateya/crcc/api/analysis"
	"github.com/ahmadateya/crcc/api/models"
	containerPkg "github.com/ahmadateya/crcc/cmd/container"
)

func applyFileSystemAnalysis(containerId string) ([]models.FileInfo, error) {
	// get container files
	containerFiles, err := containerPkg.ListContainerFilesChangesSecondVersion(containerId)
	if err != nil {
		return nil, err
	}
	// search for malicious files
	malFiles, err := analysis.FileAnalysisByNameVersionTwo(containerFiles)
	if err != nil {
		return nil, err
	}
	return malFiles, nil
}
