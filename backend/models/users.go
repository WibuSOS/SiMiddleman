package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Nama     string `json:"nama" gorm:"not null; type:varchar(255)"`
	Email    string `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password string `json:"password" gorm:"not null;type:varchar(255)"`
	NoHp     string `json:"nohp" gorm:"varchar(14)"`
	NoRek    int    `json:"norek"`
}
