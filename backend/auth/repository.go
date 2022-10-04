package auth

import (
	"gorm.io/gorm"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Repository interface {
	Login(req DataRequest) (models.Users, *errors.RestError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(req DataRequest) (models.Users, *errors.RestError) {
	var user models.Users

	res := r.db.Where("email = ? AND password = ?", req.Email, req.Password).Find(&user)

	if res.Error != nil {
		return models.Users{}, errors.NewInternalServerError("Error while fetching data")
	}

	if user.Email == "" {
		return models.Users{}, errors.NewBadRequestError("User not found")
	}

	return user, nil
}
