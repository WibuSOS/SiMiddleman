package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	RoomsID uint `gorm:"not null"`
}
