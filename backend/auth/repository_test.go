package auth

import (
	"testing"

	"github.com/WibuSOS/sinarmas/models"
	//"github.com/WibuSOS/sinarmas/utils/errors"

	"github.com/stretchr/testify/assert"
	//"gorm.io/driver/postgres"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	// config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_NAME"),
	// 	os.Getenv("DB_PORT"),
	// )
	// config := "host=localhost user=postgres password=postgres dbname=simiddleman port=5432 sslmode=disable"
	// db, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{})
	assert.NoError(t, err)

	//var user models.Users
	user := models.Users{Email: "fikri@gmail.com", Password: "fikri123"}
	result := db.Create(&user)
	assert.NoError(t, result.Error)

	return db
}

func TestLoginSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequest{
		Email:    "fikri@gmail.com",
		Password: "fikri123",
	}

	user, err := repo.Login(req)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestLoginErrorBadRequest(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequest{
		Email:    "lubis@gmail.com",
		Password: "fikri123",
	}

	_, err := repo.Login(req)
	assert.NotNil(t, err)
	assert.Equal(t, "User not found", err.Message)
	assert.Equal(t, 400, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

// func TestLoginErrorFetchingData(t *testing.T) {
// 	db := newTestDB(t)
// 	repo := NewRepository(db)

// 	req := DataRequest{
// 		Email: "fikri@gmail.com",
// 	}

// 	_, err := repo.Login(req)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "Error while fetching data", err.Message)
// 	assert.Equal(t, 500, err.Status)
// 	assert.Equal(t, "Internal_Server_Error", err.Error)
// }
