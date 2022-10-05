package models

import (
	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	PenjualID   uint `gorm:"not null"`
	PembeliID   uint
	Product     Products     `gorm:"foreignKey:RoomsID"`
	Transaction Transactions `gorm:"foreignKey:RoomsID"`
}
