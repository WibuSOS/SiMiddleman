package users

import (
	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(req *models.Users) *errors.RestError
	// GetUser() (models.Users, error)
	// UpdateUser(taskId string) error
	// DeleteUser(taskId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateUser(req *models.Users) *errors.RestError {
	pb, err := bcrypt.GenerateFromPassword([]byte(req.Password), 8)

	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	req.Role = "consumer"
	err = req.ValidateUser()
	req.Password = string(pb)
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	res := r.db.Create(&req)
	if res.Error != nil {
		return errors.NewBadRequestError(res.Error.Error())
	}

	return nil
}

// func (r *repository) GetUser() (models.Users, error) {
// 	var todos models.Users
// 	res := r.db.Find(&todos)
// 	if res.Error != nil {
// 		return models.Users{}, res.Error
// 	}

// 	return todos, nil
// }

// func (r *repository) UpdateUser(taskId string) error {
// 	idConv, _ := strconv.Atoi(taskId)
// 	todo := models.Users{}
// 	res := r.db.First(&todo, idConv)

// 	if res.Error != nil {
// 		return res.Error
// 	}

// 	if todo.Done {
// 		res = r.db.Model(&todo).Where("ID = ?", idConv).Update("Done", false)
// 	} else {
// 		res = r.db.Model(&todo).Where("ID = ?", idConv).Update("Done", true)
// 	}

// 	if res.Error != nil {
// 		return res.Error
// 	}

// 	return nil
// }

// func (r *repository) DeleteUser(taskId string) error {
// 	idConv, _ := strconv.Atoi(taskId)
// 	todo := models.Users{}

// 	res := r.db.Where("ID = ?", idConv).Delete(&todo)

// 	if res.Error != nil {
// 		return res.Error
// 	}

// 	return nil
// }
