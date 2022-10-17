package users

import (
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type Service interface {
	CreateUser(req *DataRequest) *errors.RestError
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateUser(req *DataRequest) *errors.RestError {
	err := s.repo.CreateUser(req)
	return err
}
