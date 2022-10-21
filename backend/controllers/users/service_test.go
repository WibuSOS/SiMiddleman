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

	res, err2 := service.GetUserDetails("1")
	assert.Nil(t, err2)
	assert.NotEmpty(t, res)
}

func TestGetUserDetailsError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	res, err := service.GetUserDetails("10")
	assert.NotNil(t, err)
	assert.Empty(t, res)
}

func TestUserUpdate(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	reqUpdateUser := DataRequestUpdateProfile{
		Nama:  "Ferdi Sambo",
		NoHp:  "+6281234567891",
		Email: "fsambo@gmail.com",
		NoRek: "618101052",
	}

	err := service.UpdateUser("1", reqUpdateUser)
	assert.Nil(t, err)
}

func TestServiceUpdateUserError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	reqUpdateUser := DataRequestUpdateProfile{
		Nama:  "Ferdi Sambo",
		NoHp:  "+6281234567891",
		Email: "fsambo@gmail.com",
		NoRek: "618101052",
	}

	err := service.UpdateUser("abc", reqUpdateUser)
	assert.NotNil(t, err)
}
