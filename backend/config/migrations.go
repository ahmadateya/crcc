package config

import (
	"github.com/jinzhu/gorm"
)

type Container struct {
	gorm.Model
	Name string `json:"name"`
	Scan string `json:"scan"`
}
