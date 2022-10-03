package models

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	Task string `json:"task"`
	Done bool   `json:"done"`
}
