package users

import (
	"log"
	"strconv"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	CreateUser(req *DataRequest) *errors.RestError
	GetUserDetails(userId string) (models.Users, *errors.RestError)
	UpdateUser(userId string, req DataRequest) *errors.RestError
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(req *DataRequest) *errors.RestError {
	err := req.ValidateReq()
	if err != nil {
		return err
	}

	pb, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 8)
	newUser := models.Users{
		Nama:     req.Nama,
		Role:     "consumer",
		NoHp:     req.NoHp,
		Email:    req.Email,
		Password: string(pb),
		NoRek:    req.NoRek,
	}
	res := r.db.Omit(clause.Associations).Create(&newUser)
	if res.Error != nil {
		return errors.NewBadRequestError(res.Error.Error())
	}

	return nil
}

func (r *repository) GetUserDetails(userId string) (models.Users, *errors.RestError) {
	var user models.Users

	id, _ := strconv.ParseUint(userId, 10, 64)
	res := r.db.Where("id = ?", id).Find(&user)

	if res.Error != nil {
		log.Println("Get User Details: Error while fetching data")
		return models.Users{}, errors.NewInternalServerError("Error while fetching data")
	}

	return user, nil
}

func (r *repository) UpdateUser(userId string, req DataRequest) *errors.RestError {
	idUser, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	user := models.Users{
		Nama:     req.Nama,
		NoHp:     req.NoHp,
		Email:    req.Email,
		Password: req.Password,
		NoRek:    req.NoRek,
	}

	err = r.db.Where("ID = ?", idUser).Updates(&user).First(&user).Error
	if err != nil {
		return errors.NewBadRequestError("bad request")
	}

	return nil
}
