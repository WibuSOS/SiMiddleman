package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password string `json:"password" gorm:"not null;type:varchar(255)"`
}
