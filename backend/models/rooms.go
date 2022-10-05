package models

import (
	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	PenjualID   uint `gorm:"not null"`
	PembeliID   uint
	Product     Product      `gorm:"foreignKey:RoomsID"`
	Transaction Transactions `gorm:"foreignKey:RoomsID"`
}
