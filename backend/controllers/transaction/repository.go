package transaction

import (
	"gorm.io/gorm"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Repository interface {
	GetPaymentDetails(idRoom int) (models.Rooms, *errors.RestError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetPaymentDetails(idRoom int) (models.Rooms, *errors.RestError) {
	return models.Rooms{}, nil
}
