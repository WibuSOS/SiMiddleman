package rooms

import (
	//"github.com/WibuSOS/sinarmas/backend/controllers/users"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type Service interface {
	CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError)
	GetAllRooms(userId string) ([]models.Rooms, *errors.RestError)
	JoinRoom(roomId string, userId string) (models.Rooms, *errors.RestError)
	JoinRoomPembeli(roomId string, userId string) *errors.RestError
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError) {
	newRoom, err := s.repo.CreateRoom(req)
	return newRoom, err
}

func (s *service) GetAllRooms(userId string) ([]models.Rooms, *errors.RestError) {
	newRooms, err := s.repo.GetAllRooms(userId)
	return newRooms, err
}

func (s *service) JoinRoom(roomId string, userId string) (models.Rooms, *errors.RestError) {
	room, err := s.repo.JoinRoom(roomId, userId)
	if err != nil {
		return models.Rooms{}, err
		//return DataResponse{}, err
	}

	// res := DataResponse{
	// 	ID:        room.ID,
	// 	CreatedAt: room.CreatedAt,
	// 	UpdatedAt: room.UpdatedAt,
	// 	PenjualID: room.PenjualID,
	// 	PembeliID: room.PembeliID,
	// 	RoomCode:  room.RoomCode,
	// 	Status:    room.Status,
	// 	Product:   room.Product,
	// 	Penjual: &users.DataResponse{
	// 		Nama:  room.Penjual.Nama,
	// 		NoHp:  room.Penjual.NoHp,
	// 		Email: room.Penjual.Email,
	// 		NoRek: room.Penjual.NoRek,
	// 	},
	// 	Pembeli: &users.DataResponse{
	// 		Nama:  room.Pembeli.Nama,
	// 		NoHp:  room.Pembeli.NoHp,
	// 		Email: room.Pembeli.Email,
	// 		NoRek: room.Pembeli.NoRek,
	// 	},
	// }

	return room, nil
}

func (s *service) JoinRoomPembeli(roomId string, userId string) *errors.RestError {
	err := s.repo.JoinRoomPembeli(roomId, userId)
	if err != nil {
		return err
	}
	return nil
}
