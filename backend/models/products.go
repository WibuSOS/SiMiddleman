package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	RoomsID   uint   `json:"idroom"`
	Nama      string `json:"nama" gorm:"type:varchar(100)"`
	Deskripsi string `json:"deskripsi" gorm:"type:varchar(255)"`
	Harga     uint   `json:"harga"`
	Kuantitas uint   `json:"kuantitas"`
}
