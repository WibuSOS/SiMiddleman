package users

import (
	"net/http"
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

	err = db.AutoMigrate(&models.Users{})
	// err = db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{})
	assert.NoError(t, err)

	res := db.Exec("PRAGMA foreign_keys = ON", nil)
	assert.NoError(t, res.Error)

	return db
}

func newTestDBError(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	return db
}

func TestCreateUserRepositorySuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// SUCCESS USER 1
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "1234",
	}

	err := repo.CreateUser(&req)
	assert.Empty(t, err)

	// SUCCESS USER 2
	req = DataRequest{
		Nama:     "abcde",
		NoHp:     "+6282876443890",
		Email:    "admin@abc.com",
		Password: "123456781234567812",
		NoRek:    "6789",
	}

	err = repo.CreateUser(&req)
	assert.Empty(t, err)
}

func TestCreateUserRepositoryErrorDbLevel(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// SUCCESS USER 1
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "1234",
	}

	err := repo.CreateUser(&req)
	assert.Empty(t, err)

	// ERROR USER 2
	req = DataRequest{
		Nama:     "abcde",
		NoHp:     "+6282876443890",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "6789",
	}

	err = repo.CreateUser(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

func TestCreateUserRepositoryErrorNama(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ERROR NAMA
	req := DataRequest{
		Nama:     "",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "1234",
	}

	err := repo.CreateUser(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "nama tidak memenuhi syarat", err.Message)
}

func TestCreateUserRepositoryErrorNoHp(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ERROR NO HP
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+62837",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "1234",
	}

	err := repo.CreateUser(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "no hp tidak memenuhi syarat", err.Message)
}

func TestCreateUserRepositoryErrorEmail(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ERROR EMAIL
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+6283785332789",
		Email:    "admin@",
		Password: "123456781234567812",
		NoRek:    "1234",
	}

	err := repo.CreateUser(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "email tidak memenuhi syarat", err.Message)
}

func TestCreateUserRepositoryErrorPassword(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ERROR PASSWORD
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: "1234",
		NoRek:    "1234",
	}

	err := repo.CreateUser(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "password tidak memenuhi syarat", err.Message)
}

func TestCreateUserRepositoryErrorNoRek(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ERROR NO REKENING
	req := DataRequest{
		Nama:     "vwxyz",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: "123456781234567812",
		NoRek:    "1",
	}

	err := repo.CreateUser(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "no rek tidak memenuhi syarat", err.Message)
}

func TestGetUserDetailsRepository(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequest{
		Nama:     "Julyus Andreas",
		NoHp:     "+6281234567890",
		Email:    "julyusmanurung@gmail.com",
		Password: "julyus123",
		NoRek:    "6181801052",
	}

	createUser := repo.CreateUser(&req)
	assert.Nil(t, createUser)

	userDetails, err := repo.GetUserDetails("1")
	assert.NotEmpty(t, userDetails)
	assert.Nil(t, err)
}

func TestGetUserDetailsErrorFetchingData(t *testing.T) {
	dbError := newTestDBError(t)
	db := newTestDB(t)
	repo := NewRepository(db)
	repo2 := NewRepository(dbError)

	req := DataRequest{
		Nama:     "Julyus Andreas",
		NoHp:     "+6281234567890",
		Email:    "julyusmanurung@gmail.com",
		Password: "julyus123",
		NoRek:    "6181801052",
	}

	createUser := repo.CreateUser(&req)
	assert.Nil(t, createUser)

	userDetails, err := repo2.GetUserDetails("1")
	assert.Empty(t, userDetails)
	assert.NotNil(t, err)
	assert.Equal(t, "Error while fetching data", err.Message)
}
