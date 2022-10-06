package rooms

import (
	"math/rand"
	"time"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError)
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

func generateRoomCode(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

func (r *repository) CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError) {
	if err := req.ValidateReq(); err != nil {
		return models.Rooms{}, err
	}

	roomCodeLength := 10
	newRoom := models.Rooms{
		PenjualID: req.PenjualID,
		RoomCode:  generateRoomCode(roomCodeLength),
		Product: &models.Products{
			Nama:      req.Product.Nama,
			Deskripsi: req.Product.Deskripsi,
			Harga:     req.Product.Harga,
			Kuantitas: req.Product.Kuantitas,
		},
	}

	res := r.db.Omit("Transaction").Create(&newRoom)
	if res.Error != nil {
		return models.Rooms{}, errors.NewBadRequestError(res.Error.Error())
	}

	return newRoom, nil
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
