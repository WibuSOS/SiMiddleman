package rooms

import (
	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Service interface {
	CreateRoom(req *DataRequest) (models.Rooms, *errors.RestError)
	GetAllRooms(user_id string) ([]models.Rooms, *errors.RestError)
	// UpdateUser(taskId string) (int, error)
	// DeleteUser(taskId string) (int, error)
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

// func (s *service) UpdateUser(taskId string) (int, error) {
// 	err := s.repo.UpdateUser(taskId)
// 	if err != nil {
// 		return http.StatusInternalServerError, err
// 	}

// 	return http.StatusOK, nil
// }

// func (s *service) DeleteUser(taskId string) (int, error) {
// 	err := s.repo.DeleteUser(taskId)
// 	if err != nil {
// 		return http.StatusInternalServerError, err
// 	}

// 	return http.StatusOK, nil
// }
