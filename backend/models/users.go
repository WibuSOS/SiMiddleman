package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Nama     string `json:"nama" gorm:"not null;type:varchar(30)"`
	Role     string `json:"role" gorm:"not null;type:varchar(15);default:consumer"`
	NoHp     string `json:"noHp" gorm:"type:varchar(18)"`
	Email    string `json:"email" gorm:"unique;not null;type:varchar(30)"`
	Password string `json:"password" gorm:"not null;type:varchar(128)"`
	NoRek    string `json:"noRek" gorm:"type:varchar(18)"`
}
