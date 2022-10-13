package transaction

import (
	"log"
	"strconv"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateStatusDelivery(id string, req RequestUpdateStatus) *errors.RestError
	GetPaymentDetails(idRoom int) (models.Rooms, *errors.RestError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateStatusDelivery(id string, req RequestUpdateStatus) *errors.RestError {
	idroom, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err.Error())
	}

	room := models.Rooms{
		Model: gorm.Model{
			ID: uint(idroom),
		},
		Status: req.Status,
	}
	err = r.db.Model(&room).Updates(&room).Error

	if err != nil {
		log.Println("Update Status error : ", err.Error())
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
