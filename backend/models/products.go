package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	RoomsID   uint   `json:"idroom" gorm:"not null"`
	Nama      string `json:"nama"`
	Harga     int    `json:"harga"`
	Kuantitas int    `json:"kuantitas"`
	Deskripsi string `json:"deskripsi"`
}
