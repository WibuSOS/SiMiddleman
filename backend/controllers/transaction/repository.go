package transaction

import (
	"log"
	"strconv"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateStatusDelivery(id string) *errors.RestError
	GetPaymentDetails(idRoom int) (models.Rooms, *errors.RestError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateStatusDelivery(id string) *errors.RestError {
	idroom, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err.Error())
	}

	room := models.Rooms{
		Model: gorm.Model{
			ID: uint(idroom),
		},
		Status: "barang dikirim",
	}
	err = r.db.Model(&room).Updates(&room).Error

	if err != nil {
		log.Println("Update Status Barang error : ", err.Error())
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (r *repository) GetPaymentDetails(idRoom int) (models.Rooms, *errors.RestError) {
	return models.Rooms{}, nil
}
