package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	RoomsID   uint   `json:"roomsID,omitempty"`
	Nama      string `json:"nama,omitempty" gorm:"type:varchar(100)"`
	Deskripsi string `json:"deskripsi,omitempty" gorm:"type:varchar(255)"`
	Harga     uint   `json:"harga,omitempty"`
	Kuantitas uint   `json:"kuantitas,omitempty"`
}
