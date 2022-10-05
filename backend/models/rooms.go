package models

import (
	"gorm.io/gorm"
)

type Rooms struct {
	gorm.Model
	PenjualID   uint
	PembeliID   *uint
	RoomCode    string       `gorm:"not null;unique;type:varchar(15)"`
	Product     Products     `gorm:"foreignKey:RoomsID"`
	Transaction Transactions `gorm:"foreignKey:RoomsID"`
}
