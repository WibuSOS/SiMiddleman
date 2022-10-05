package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	IdRoom    int    `json:"idroom"`
	Nama      string `json:"nama"`
	Harga     int    `json:"harga"`
	Kuantitas int    `json:"kuantitas"`
	Deskripsi string `json:"deskripsi"`
}
