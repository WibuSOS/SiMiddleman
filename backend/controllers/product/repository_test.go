package product

import (
	"testing"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{}, &models.Products{}, &models.Rooms{})
	assert.NoError(t, err)

	db.Create(&models.Users{
		Nama:     "test2",
		NoHp:     "081234523415",
		NoRek:    "12341415",
		Email:    "test12@gmail.com",
		Password: "123456789",
	})

	db.Create(&models.Products{
		Nama:      "ayam",
		Harga:     15000,
		Kuantitas: 2,
		Deskripsi: "ini adalah ayam",
	})

	return db
}

func TestRepoUpdateProductSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequest{
		Nama:      "hello",
		Deskripsi: "bandung",
		Harga:     230000,
		Kuantitas: 2,
	}

	update, err := repo.UpdateProduct("1", req)
	assert.Empty(t, err)
	assert.NotEmpty(t, update)
}
