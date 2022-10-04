package auth

import (

	//"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
	"github.com/WibuSOS/sinarmas/utils/token"
)

type Service interface {
	Login(req DataRequest) (DataResponse, *errors.RestError)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Login(req DataRequest) (DataResponse, *errors.RestError) {
	if err := req.Validation(); err != nil {
		return DataResponse{}, err
	}

	user, err := s.repo.Login(req)
	if err != nil {
		return DataResponse{}, err
	}

	token, err := token.GenerateToken(user)
	if err != nil {
		return DataResponse{}, err
	}

	res := DataResponse{
		Email: user.Email,
		Token: token,
	}

	return res, nil
}
