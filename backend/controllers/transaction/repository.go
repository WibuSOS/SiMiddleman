package transaction

import (
	"log"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateStatusDelivery(id string) *errors.RestError
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateStatusDelivery(id string) *errors.RestError {
	room := models.Rooms{
		Status: "barang dikirim",
	}
	err := r.db.Where("id = ?", id).Updates(&room).Error

	if err != nil {
		log.Println("Update Status Barang error : ", err.Error())
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}
