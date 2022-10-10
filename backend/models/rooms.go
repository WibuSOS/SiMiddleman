package models

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	PenjualID   uint          `json:"penjualID,omitempty"`
	PembeliID   *uint         `json:"pembeliID,omitempty"`
	RoomCode    string        `json:"roomCode,omitempty" gorm:"not null;unique;type:varchar(15)"`
	Status      string        `json:"status,omitempty" gorm:"varchar(20);default:mulai transaksi"`
	Product     *Products     `json:"product,omitempty" gorm:"foreignKey:RoomsID"`
	Transaction *Transactions `json:"transaction,omitempty" gorm:"foreignKey:RoomsID"`
}
