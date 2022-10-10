package transaction

import (
	"log"

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

func (r *repository) GetPaymentDetails(idRoom int) (models.Rooms, *errors.RestError) {
	var room models.Rooms

	res := r.db.Where("id = ?", idRoom).Preload("Product").Find(&room)
	if res.Error != nil {
		log.Println("Get Payment Details: Error while fetching data")
		return models.Rooms{}, errors.NewInternalServerError("Error while fetching data")
	}

	if room.RoomCode == "" {
		log.Println("Get Payment Details: Room not found")
		return models.Rooms{}, errors.NewBadRequestError("Room not found")
	}

	return room, nil
}
