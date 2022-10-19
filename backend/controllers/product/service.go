package product

import (
	//"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type Service interface {
	UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError) {
	if err := req.ValidateReq(); err != nil {
		return models.Products{}, err
	}
	product, err := s.repo.UpdateProduct(id, req)
	if err != nil {
		return models.Products{}, err
	}

	return product, nil
}
