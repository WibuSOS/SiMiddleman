package rooms

import (
	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Service interface {
	CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError)
	GetAllRooms(user_id string) ([]models.Rooms, *errors.RestError)
	JoinRoom(room_id string, user_id string) (models.Rooms, *errors.RestError)
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

func (s *service) GetAllRooms(user_id string) ([]models.Rooms, *errors.RestError) {
	newRooms, err := s.repo.GetAllRooms(user_id)
	return newRooms, err
}

func (s *service) JoinRoom(room_id string, user_id string) (models.Rooms, *errors.RestError) {
	room, err := s.repo.JoinRoom(room_id, user_id)
	if err != nil {
		return models.Rooms{}, err
	}
	return room, nil
}
