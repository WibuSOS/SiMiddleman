package product

import (
	//"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Service interface {
	GetSpesifikProduct(idroom string, req DataRequest) (models.Products, *errors.RestError)
	CreateProduct(idroom string, req DataRequest) (models.Products, *errors.RestError)
	UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError)
	DeleteProduct(id string) *errors.RestError
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetSpesifikProduct(idroom string, req DataRequest) (models.Products, *errors.RestError) {
	product, err := s.repo.GetSpesifikProduct(idroom)
	if err != nil {
		return models.Products{}, err
	}

	return product, nil
}

func (s *service) CreateProduct(idroom string, req DataRequest) (models.Products, *errors.RestError) {
	if err := req.Validation(); err != nil {
		return models.Products{}, err
	}
	product, err := s.repo.CreateProduct(idroom, req)
	if err != nil {
		return models.Products{}, err
	}

	return product, nil
}

func (s *service) UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError) {
	product, err := s.repo.UpdateProduct(id, req)
	if err != nil {
		return models.Products{}, err
	}

	return product, nil
}

func (s *service) DeleteProduct(id string) *errors.RestError {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
