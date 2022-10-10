package transaction

import (

	//"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Service interface {
	GetPaymentDetails(idRoom int) (ResponsePaymentInfo, *errors.RestError)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetPaymentDetails(idRoom int) (ResponsePaymentInfo, *errors.RestError) {
	return ResponsePaymentInfo{}, nil
}
