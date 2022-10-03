package todos

import (
	"github.com/WibuSOS/sinarmas/models"
	"net/http"
)

type Service interface {
	GetTodos() ([]models.Todos, int, error)
	CreateTodos(req DataRequest) (models.Todos, int, error)
	CheckTodo(taskId string) (int, error)
	DeleteTodo(taskId string) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetTodos() ([]models.Todos, int, error) {
	todos, err := s.repo.GetTodos()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return todos, http.StatusOK, nil
}

func (s *service) CreateTodos(req DataRequest) (models.Todos, int, error) {
	todo, err := s.repo.CreateTodos(req.Task)
	if err != nil {
		return models.Todos{}, http.StatusInternalServerError, err
	}

	return todo, http.StatusOK, nil
}

func (s *service) CheckTodo(taskId string) (int, error) {
	err := s.repo.CheckTodo(taskId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *service) DeleteTodo(taskId string) (int, error) {
	err := s.repo.DeleteTodo(taskId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
