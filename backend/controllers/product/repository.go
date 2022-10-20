package product

import (
	"log"

	"gorm.io/gorm"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"
)

type Repository interface {
	UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError) {
	product := models.Products{
		Nama:      req.Nama,
		Harga:     req.Harga,
		Kuantitas: req.Kuantitas,
		Deskripsi: req.Deskripsi,
	}
	r.db.Where("id = ?", id).Updates(&product)

	err := r.db.First(&product, "id = ?", id).Find(&product).Error
	if err != nil {
		log.Println("Get The Update Data error : ", err.Error())
		return models.Products{}, errors.NewBadRequestError(err.Error())
	}

	return product, nil
}
