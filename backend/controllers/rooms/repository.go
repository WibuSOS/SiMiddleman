package rooms

import (
	"log"
	"strconv"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
	"github.com/WibuSOS/sinarmas/backend/utils/localization"

	"github.com/dchest/uniuri"
	"gorm.io/gorm"
)

type Repository interface {
	CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError)
	GetAllRooms(userId string) ([]models.Rooms, *errors.RestError)
	JoinRoom(roomId string, userId string) (models.Rooms, *errors.RestError)
	JoinRoomPembeli(roomId string, userId string, mesasge string) *errors.RestError
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func GenerateRoomCode(length int, roomId uint) string {
	s := uniuri.NewLen(length) + strconv.FormatUint(uint64(roomId), 10)

	return s
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
		res := tx.Omit("Transaction", "Penjual", "Pembeli").Create(&newRoom)
		if res.Error != nil {
			return res.Error
		}

		res = tx.Model(&newRoom).Omit("Transaction", "Penjual", "Pembeli").Update("RoomCode", GenerateRoomCode(roomCodeLength, newRoom.ID))
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

func (r *repository) GetAllRooms(userId string) ([]models.Rooms, *errors.RestError) {
	var user models.Users
	var newRooms []models.Rooms

	id, _ := strconv.ParseUint(userId, 10, 64)
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

func (r *repository) JoinRoom(roomId string, userId string) (models.Rooms, *errors.RestError) {
	var room models.Rooms

	res := r.db.
		Preload("Product").
		Preload("Penjual").
		Preload("Pembeli").
		Where("id = ? AND (penjual_id = ? OR pembeli_id = ?)", roomId, userId, userId).
		Find(&room)

	if res.Error != nil {
		log.Println(res.Error.Error())
		return models.Rooms{}, errors.NewInternalServerError("internalServer")
	}

	if room.ID == 0 {
		return models.Rooms{}, errors.NewBadRequestError("tidakBisaMasukRuangan")
	}

	return room, nil
}

func (r *repository) JoinRoomPembeli(roomId string, userId string, message string) *errors.RestError {
	var room models.Rooms

	idroom64, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		log.Println(err.Error())
		msg := localization.GetMessage(message, "invalididuser")
		return errors.NewBadRequestError(msg)
	}
	idRoom := uint(idroom64)

	alreadyJoinRoom := r.db.
		Where("room_code = ? AND (penjual_id = ? OR pembeli_id = ?)", roomId, idRoom, idRoom).
		First(&room)
	if alreadyJoinRoom.Error == nil {
		msg := localization.GetMessage(message, "sudahmasukruangan")
		return errors.NewBadRequestError(msg)
	}

	roomAlreadyHasPembeli := r.db.
		Where("room_code = ? AND pembeli_id IS NULL", roomId).
		First(&room)
	if roomAlreadyHasPembeli.Error != nil {
		msg := localization.GetMessage(message, "sudahadapembeli")
		return errors.NewBadRequestError(msg)
	}

	res := r.db.
		Where("room_code = ? AND pembeli_id IS NULL", roomId).
		Updates(models.Rooms{
			PembeliID: &idRoom,
		})
	if res.Error != nil {
		return errors.NewBadRequestError(res.Error.Error())
	}

	return nil
}
