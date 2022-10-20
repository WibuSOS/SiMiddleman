package auth

import (
	"testing"

	"github.com/WibuSOS/sinarmas/backend/database"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := database.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	return db
}

func TestLoginSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequest{
		Email:    "penjual@custom.com",
		Password: "12345678",
	}

	user, err := repo.Login(req)
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestLoginErrorUserNotFound(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequest{
		Email:    "lubis@gmail.com",
		Password: "fikri123",
	}

	_, err := repo.Login(req)
	assert.NotNil(t, err)
	assert.Equal(t, "userNotFound", err.Message)
	assert.Equal(t, 400, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

func TestLoginErrorAuthenticationFailed(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequest{
		Email:    "penjual@custom.com",
		Password: "lubis123",
	}

	_, err := repo.Login(req)
	assert.NotNil(t, err)
	assert.Equal(t, "authFailed", err.Message)
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
