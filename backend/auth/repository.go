package auth

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/WibuSOS/SiMiddleman/models"
	//"github.com/WibuSOS/SiMiddleman/utils/errors"
)

type Repository interface {
	Login(req DataRequest) (models.Users, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(req DataRequest) (models.Users, error) {
	var user models.Users

	res := r.db.Where("email = ? AND password = ?", req.Email, req.Password).Find(&user)
	//fmt.Println(user, "user luar")
	if res.Error != nil || user.Email == "" {
		fmt.Println(user, "user dalam")
		return models.Users{}, res.Error
	}

	return user, nil
}
