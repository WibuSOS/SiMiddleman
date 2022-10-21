package users

import (
	"log"

	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type Service interface {
	CreateUser(req *DataRequest) *errors.RestError
	GetUserDetails(idUser string) (DataResponse, *errors.RestError)
	UpdateUser(idUser string, req DataRequestUpdateProfile) *errors.RestError
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

func (s *service) GetUserDetails(idUser string) (DataResponse, *errors.RestError) {
	user, err := s.repo.GetUserDetails(idUser)
	if err != nil {
		log.Println(err.Message)
		return DataResponse{}, err
	}

	res := DataResponse{
		Nama:  user.Nama,
		NoHp:  user.NoHp,
		Email: user.Email,
		NoRek: user.NoRek,
	}

	return res, nil
}

func (s *service) UpdateUser(idUser string, req DataRequestUpdateProfile) *errors.RestError {

	err := s.repo.UpdateUser(idUser, req)
	if err != nil {
		log.Println(err.Message)
		return err
	}

	return nil
}
