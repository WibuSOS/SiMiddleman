package rooms

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError)
	GetAllRooms(user_id string) ([]models.Rooms, *errors.RestError)
	JoinRoom(room_id string, user_id string) (models.Rooms, *errors.RestError)
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

func (r *repository) GetAllRooms(user_id string) ([]models.Rooms, *errors.RestError) {
	var user models.Users
	var newRooms []models.Rooms

	id, _ := strconv.ParseUint(user_id, 10, 64)
	res := r.db.
		Preload("PenjualRooms.Product").
		Preload("PenjualRooms.Transaction").
		Preload("PembeliRooms.Product").
		Preload("PembeliRooms.Transaction").
		First(&user, id)
	if res.Error != nil {
		return []models.Rooms{}, errors.NewBadRequestError(res.Error.Error())
	}

	newRooms = append(newRooms, user.PenjualRooms...)
	newRooms = append(newRooms, user.PembeliRooms...)

	return newRooms, nil
}

func (r *repository) JoinRoom(room_id string, user_id string) (models.Rooms, *errors.RestError) {
	var room models.Rooms

	res := r.db.
		Preload("Product").
		Where("id = ? AND (penjual_id = ? OR pembeli_id = ?)", room_id, user_id, user_id).Find(&room)

	if res.Error != nil {
		return models.Rooms{}, errors.NewBadRequestError(res.Error.Error())
	}

	if room.ID == 0 {
		return models.Rooms{}, errors.NewBadRequestError("Tidak bisa memasuki ruangan")
	}

	return room, nil
}

// func (r *repository) DeleteUser(taskId string) error {
// 	idConv, _ := strconv.Atoi(taskId)
// 	todo := models.Users{}

// 	res := r.db.Where("ID = ?", idConv).Delete(&todo)

// 	if res.Error != nil {
// 		return res.Error
// 	}

// 	return nil
// }
