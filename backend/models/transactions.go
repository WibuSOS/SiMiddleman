package models

import "gorm.io/gorm"

type Transactions struct {
	gorm.Model
	RoomsID uint `json:"roomsID,omitempty"`
}
