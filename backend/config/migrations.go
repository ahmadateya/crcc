package config

type Container struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `json:"name"`
	Scan      string `json:"scan"`
	CreatedAt string `json:"CreatedAt"`
}
