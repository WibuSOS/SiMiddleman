package users

import (
	"github.com/WibuSOS/sinarmas/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	// GetUser() (models.Users, error)
	CreateUser(req *models.Users) error
	// UpdateUser(taskId string) error
	// DeleteUser(taskId string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) GetUser() (models.Users, error) {
// 	var todos models.Users
// 	res := r.db.Find(&todos)
// 	if res.Error != nil {
// 		return models.Users{}, res.Error
// 	}

// 	return todos, nil
// }

func (r *repository) CreateUser(req *models.Users) error {
	pb, err := bcrypt.GenerateFromPassword([]byte(req.Password), 4)

	if err != nil {
		return err
	}

	req.Role = "consumer"
	req.Password = string(pb)
	res := r.db.Create(&req)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

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
