package rooms

import (
	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRoom(req *DataRequest) *errors.RestError
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

func (r *repository) CreateRoom(req *DataRequest) *errors.RestError {
	if err := req.ValidateReq(); err != nil {
		return err
	}

	newRoom := models.Rooms{
		PenjualID: req.PenjualID,
		Product:   models.Product{},
	}

	res := r.db.Omit("Transaction").Create(&newRoom)
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
