package auth

import (
	"testing"

	//"github.com/WibuSOS/sinarmas/models"
	//"github.com/WibuSOS/sinarmas/utils/errors"

	"github.com/stretchr/testify/assert"
)

func TestLoginServiceSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Email:    "fikri@gmail.com",
		Password: "fikri123",
	}

	res, token, err := service.Login(req)
	assert.Nil(t, err)
	assert.NotEqual(t, "", res.Email)
	assert.NotEqual(t, "", token)
}

func TestLoginServiceErrorInvalidEmail(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Email:    "",
		Password: "fikri123",
	}

	_, _, err := service.Login(req)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid email", err.Message)
	assert.Equal(t, 400, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

func TestLoginServiceErrorInvalidPassword(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Email:    "fikri@gmail.com",
		Password: "",
	}

	_, _, err := service.Login(req)
	assert.NotNil(t, err)
	assert.Equal(t, "Invalid password", err.Message)
	assert.Equal(t, 400, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

func TestLoginServiceErrorUserNotFound(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Email:    "lubis@gmail.com",
		Password: "fikri123",
	}

	_, _, err := service.Login(req)
	assert.NotNil(t, err)
	assert.Equal(t, "User not found", err.Message)
	assert.Equal(t, 400, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}
