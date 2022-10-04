package users

import (
	"net/http"

	"github.com/WibuSOS/sinarmas/models"
)

type Service interface {
	// GetUser() (models.Users, int, error)
	CreateUser(req *models.Users) (int, error)
	// UpdateUser(taskId string) (int, error)
	// DeleteUser(taskId string) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

// func (s *service) GetUser() (models.Users, int, error) {
// 	todos, err := s.repo.GetTodos()
// 	if err != nil {
// 		return models.Users{}, http.StatusInternalServerError, err
// 	}

// 	return todos, http.StatusOK, nil
// }

func (s *service) CreateUser(req *models.Users) (int, error) {
	err := s.repo.CreateUser(req)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
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
