package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
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

	res := r.db.Where("email = ?", req.Email).Find(&user)

	if res.Error != nil {
		log.Println(res.Error)
		return models.Users{}, errors.NewInternalServerError("internalServer")
	}

	if user.Email == "" {
		return models.Users{}, errors.NewBadRequestError("userNotFound")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return models.Users{}, errors.NewBadRequestError("authFailed")
	}

	return user, nil
}
