package models

type ListContainer struct {
	ID     string   `json:"id"`
	Names  []string `json:"names"`
	Image  string   `json:"image"`
	Status string   `json:"status"`
	Ports  []Port   `json:"ports"`
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


type FileInfo struct{
	Path string `json:"path"`
	Kind string `json:"kind"`
}

type ContainerFile struct {
	File string `json:"file"`
	Description string `json:"description"`
	Impact string `json:"impact"`
}


type ContainerProcesses struct{
	Processes [][]string
}


type ContainerProcess struct {
	Cmd string `json:"cmd"`
	User string `json:"user"`
	Type int `json:"type"`
	Description string `json:"description"`
	Impact string `json:"impact"`
}

type ContainerPorts struct {
	Port string `json:"port"`
	Description string `json:"description"`
	Impact string `json:"impact"`
}