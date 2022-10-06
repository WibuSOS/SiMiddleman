package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
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

	res := r.db.Where("email = ?", req.Email).Find(&user)

	if res.Error != nil {
		log.Println("Login: Error while fetching data")
		return models.Users{}, errors.NewInternalServerError("Error while fetching data")
	}

	if user.Email == "" {
		log.Println("Login: User not found")
		return models.Users{}, errors.NewBadRequestError("User not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println("Login: Authentication failed")
		return models.Users{}, errors.NewBadRequestError("Authentication failed")
	}

	return user, nil
}
