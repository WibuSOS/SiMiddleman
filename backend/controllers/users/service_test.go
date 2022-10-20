package users

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserServiceSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	// SUCCESS USER 1
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "1234",
	}

	err := service.CreateUser(&req)
	assert.Empty(t, err)

	// SUCCESS USER 2
	req = DataRequest{
		Nama:     "abcde",
		NoHp:     "+6282876443890",
		Email:    "admin@abc.com",
		Password: "123456781234567812",
		NoRek:    "6789",
	}

	err = service.CreateUser(&req)
	assert.Empty(t, err)
}

func TestCreateUserServiceError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	// SUCCESS USER 1
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "1234",
	}

	err := service.CreateUser(&req)
	assert.Empty(t, err)

	// ERROR USER 2
	req = DataRequest{
		Nama:     "abcde",
		NoHp:     "+6282876443890",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "6789",
	}

	err = service.CreateUser(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

func TestGetUserDetails(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := DataRequest{
		Nama:     "Julyus Andreas",
		NoHp:     "+6281234567890",
		Email:    "julyusmanurung@gmail.com",
		Password: "julyus123",
		NoRek:    "1",
	}

	err := service.CreateUser(&req)
	assert.Empty(t, err)

	res, err2 := service.GetUserDetails("1")
	assert.Nil(t, err2)
	assert.NotEmpty(t, res)
}

func TestGetUserDetailsError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	dbError := newTestDBError(t)
	repoError := NewRepository(dbError)
	serviceError := NewService(repoError)

	req := DataRequest{
		Nama:     "Julyus Andreas",
		NoHp:     "+6281234567890",
		Email:    "julyusmanurung@gmail.com",
		Password: "julyus123",
		NoRek:    "6181801052",
	}

	err := service.CreateUser(&req)
	assert.Empty(t, err)

	res, err2 := serviceError.GetUserDetails("1")
	assert.NotNil(t, err2)
	assert.Empty(t, res)
}
