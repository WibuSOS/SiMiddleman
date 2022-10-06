package product

import (
	//"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Service interface {
	// GetSpesifikProduct(idroom uint, req DataRequest) (models.Products, error)
	// CreateProduct(idroom uint, req DataRequest) (models.Products, error)
	// CreateProductReturnID(idroom uint, req DataRequest) (uint, error)
	UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError)
	DeleteProduct(id string) *errors.RestError
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

// func (s *service) GetSpesifikProduct(idroom uint, req DataRequest) (models.Products, error) {
// 	product, err := s.repo.GetSpesifikProduct(idroom)
// 	if err != nil {
// 		return models.Products{}, err
// 	}

// 	return product, nil
// }

// func (s *service) CreateProduct(idroom uint, req DataRequest) (models.Products, error) {

// 	product, err := s.repo.CreateProduct(idroom, req)
// 	if err != nil {
// 		return models.Products{}, err
// 	}

// 	return product, nil
// }

// func (s *service) CreateProductReturnID(idroom uint, req DataRequest) (uint, error) {

// 	idproduct, err := s.repo.CreateProductReturnID(idroom, req)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return idproduct, nil
// }

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

func (s *service) DeleteProduct(id string) *errors.RestError {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
