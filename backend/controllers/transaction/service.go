package transaction

import "github.com/WibuSOS/sinarmas/backend/utils/errors"

type Service interface {
	UpdateStatusDelivery(id string, req RequestUpdateStatus) *errors.RestError
	GetPaymentDetails(idRoom int) (ResponsePaymentInfo, *errors.RestError)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) UpdateStatusDelivery(id string, req RequestUpdateStatus) *errors.RestError {
	if err := req.Validation(); err != nil {
		return err
	}

	err := s.repo.UpdateStatusDelivery(id, req)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetPaymentDetails(idRoom int) (ResponsePaymentInfo, *errors.RestError) {
	room, err := s.repo.GetPaymentDetails(idRoom)
	if err != nil {
		return ResponsePaymentInfo{}, err
	}

	status := room.Status
	total := int(room.Product.Harga) * int(room.Product.Kuantitas)

	res := ResponsePaymentInfo{
		Total:  uint(total),
		Status: status,
	}

	return res, nil
}
