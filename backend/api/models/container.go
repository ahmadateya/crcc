package models

type ListContainer struct {
	ID              string      `json:"id"`
	Names           []string    `json:"names"`
	Image           string      `json:"image"`
	Status          string      `json:"status"`
	Ports           []Port      `json:"ports"`
	NetworkSettings interface{} `json:"networksettings"`
}

type Port struct {
	IP          string `json:"ip"`
	PrivatePort string `json:"private_port"`
	PublicPort  int    `json:"public_port"`
	Type        string `json:"type"`
}

type Container struct {
	ID      string `json:"id"`
	Image   string `json:"image"`
	Created string `json:"created"`
	Name    string `json:"name"`
}

type FileInfo struct {
	Path string `json:"path"`
	Kind int    `json:"Kind"`
}

type ContainerProcesses struct {
	Processes [][]string
}

type ContainerPorts struct {
	Port        string `json:"port"`
	Description string `json:"description"`
	Impact      string `json:"impact"`
}
