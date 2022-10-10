package transaction

import "github.com/WibuSOS/sinarmas/utils/errors"

type Service interface {
	UpdateStatusDelivery(id string) *errors.RestError
	GetPaymentDetails(idRoom int) (ResponsePaymentInfo, *errors.RestError)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) UpdateStatusDelivery(id string) *errors.RestError {
	err := s.repo.UpdateStatusDelivery(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetPaymentDetails(idRoom int) (ResponsePaymentInfo, *errors.RestError) {
	return ResponsePaymentInfo{}, nil
}
