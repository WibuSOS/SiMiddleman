package rooms

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError)
	GetAllRooms(user_id string) ([]models.Rooms, *errors.RestError)
	JoinRoom(room_id string, user_id string) (models.Rooms, *errors.RestError)
	JoinRoomPembeli(room_id string, user_id string) *errors.RestError
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func generateRoomCode(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
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
		Status:    "mulai transaksi",
		Product: &models.Products{
			Nama:      req.Product.Nama,
			Deskripsi: req.Product.Deskripsi,
			Harga:     req.Product.Harga,
			Kuantitas: req.Product.Kuantitas,
		},
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// do some database operations in the transaction (use 'tx' from this point, not 'db')
		res := tx.Omit("Transaction").Create(&newRoom)
		if res.Error != nil {
			return res.Error
		}

		res = tx.Model(&newRoom).Omit("Transaction").Update("RoomCode", generateRoomCode(roomCodeLength)+strconv.FormatUint(uint64(newRoom.ID), 10))
		if res.Error != nil {
			return res.Error
		}

		// return nil will commit the whole transaction
		return nil
	})

	if err != nil {
		return models.Rooms{}, errors.NewBadRequestError(err.Error())
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
		Where("id = ? AND (penjual_id = ? OR pembeli_id = ?)", room_id, user_id, user_id).
		Find(&room)

	if res.Error != nil {
		return models.Rooms{}, errors.NewBadRequestError(res.Error.Error())
	}

	if room.ID == 0 {
		return models.Rooms{}, errors.NewBadRequestError("Tidak bisa memasuki ruangan")
	}

	return room, nil
}

func (r *repository) JoinRoomPembeli(room_id string, user_id string) *errors.RestError {
	var room models.Rooms

	idroom64, err := strconv.ParseUint(user_id, 10, 32)
	if err != nil {
		log.Println(err.Error())
		return errors.NewBadRequestError("Invalid User ID")
	}
	idRoom := uint(idroom64)

	alreadyJoinRoom := r.db.
		Where("room_code = ? AND (penjual_id = ? OR pembeli_id = ?)", room_id, idRoom, idRoom).
		First(&room)
	if alreadyJoinRoom.Error == nil {
		return errors.NewBadRequestError("Anda sudah masuk kedalam room ini")
	}

	roomAlreadyHasPembeli := r.db.
		Where("room_code = ? AND pembeli_id IS NULL", room_id).
		First(&room)
	if roomAlreadyHasPembeli.Error != nil {
		return errors.NewBadRequestError(roomAlreadyHasPembeli.Error.Error())
	}

	res := r.db.
		Where("room_code = ? AND pembeli_id IS NULL", room_id).
		Updates(models.Rooms{
			PembeliID: &idRoom,
		})
	if res.Error != nil {
		return errors.NewBadRequestError(res.Error.Error())
	}

	return nil
}
