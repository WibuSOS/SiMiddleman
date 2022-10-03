package auth

import (
	"net/http"

	//"github.com/WibuSOS/SiMiddleman/models"
	//"github.com/WibuSOS/SiMiddleman/utils/errors"
	"github.com/WibuSOS/SiMiddleman/utils/token"
)

type Service interface {
	Login(req DataRequest) (DataResponse, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Login(req DataRequest) (DataResponse, int, error) {
	// if err := req.Validation(); err != nil {
	// 	return DataResponse{}, http.BadRequest,
	// }

	user, err := s.repo.Login(req)
	if err != nil || user.Email == "" {
		return DataResponse{}, http.StatusInternalServerError, err
	}

	token, _, err := token.GenerateToken(user)
	if err != nil {
		return DataResponse{}, http.StatusInternalServerError, err
	}

	res := DataResponse{
		Email: user.Email,
		Token: token,
	}

	return res, http.StatusOK, nil
}
