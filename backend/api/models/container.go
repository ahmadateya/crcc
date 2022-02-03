package models

type Container struct {
	ID     string   `json:"id"`
	Names  []string `json:"names"`
	Image  string   `json:"image"`
	Status string   `json:"status"`
	Ports  []Port   `json:"ports"`
}

type Port struct {
	IP          string `json:"ip"`
	PrivatePort string `json:"private_port"`
	PublicPort  int    `json:"public_port"`
	Type        string `json:"type"`
}
