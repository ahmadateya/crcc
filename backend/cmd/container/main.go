package container

import (
	"fmt"
	"log"
	"os/exec"
)

// container package to handle container related business

// TODO: make a connection to docker daemon on the host machine

func ListContainers() []byte {
	output, err := exec.Command("docker", "ps --format '{\"ID\":\"{{ .ID }}\", \"Image\": \"{{ .Image }}\", \"Names\":\"{{ .Names }}\"}'").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", output)
	return output
}
