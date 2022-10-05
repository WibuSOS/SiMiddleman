package product

import (
	"log"
	"strconv"

	"gorm.io/gorm"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/WibuSOS/sinarmas/utils/errors"
)

type Repository interface {
	GetSpesifikProduct(idroom string) (models.Products, *errors.RestError)
	CreateProduct(idroom string, req DataRequest) (models.Products, *errors.RestError)
	UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError)
	DeleteProduct(id string) *errors.RestError
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetSpesifikProduct(idroom string) (models.Products, *errors.RestError) {
	var product models.Products

	res := r.db.Where("id_room = ?", idroom).Find(&product)

	if res.Error != nil {
		return models.Products{}, errors.NewInternalServerError("Error while fetching data")
	}

	if product.Nama == "" {
		return models.Products{}, errors.NewBadRequestError("Belum ada Product")
	}

	return product, nil
}

func (r *repository) CreateProduct(idroom string, req DataRequest) (models.Products, *errors.RestError) {
	idroomconv, _ := strconv.Atoi(idroom)

	product := models.Products{
		IdRoom:    idroomconv,
		Nama:      req.Nama,
		Harga:     req.Harga,
		Kuantitas: req.Kuantitas,
		Deskripsi: req.Deskripsi,
	}
	res := r.db.Create(&product)
	if res.Error != nil {
		log.Println("Create Data error : ", res.Error)
		return models.Products{}, errors.NewBadRequestError(res.Error.Error())
	}

	return product, nil
}

func (r *repository) UpdateProduct(id string, req DataRequest) (models.Products, *errors.RestError) {
	product := models.Products{
		Nama:      req.Nama,
		Harga:     req.Harga,
		Kuantitas: req.Kuantitas,
		Deskripsi: req.Deskripsi,
	}
	res := r.db.Where("id = ?", id).Updates(&product)
	if res.Error != nil {
		log.Println("Update Data error : ", res.Error)
		return models.Products{}, errors.NewBadRequestError(res.Error.Error())
	}

	err := r.db.First(&product, "id = ?", id).Find(&product).Error
	if err != nil {
		log.Println("Get The Update Data error : ", res.Error)
		return models.Products{}, errors.NewBadRequestError(res.Error.Error())
	}

	return product, nil
}

func (r *repository) DeleteProduct(id string) *errors.RestError {
	product := models.Products{}
	res := r.db.Where("ID = ?", id).Delete(&product)
	if res.Error != nil {
		log.Println("Delete Data error : ", res.Error)
		return errors.NewBadRequestError(res.Error.Error())
	}

	return nil
}
