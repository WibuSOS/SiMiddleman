package users

import (
	"net/http"
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
	assert.Equal(t, "nameDoesNotQualify", err.Message)
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
	assert.Equal(t, "hpDoesNotQualify", err.Message)
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
	assert.Equal(t, "emailDoesNotQualify", err.Message)
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
	assert.Equal(t, "passwordDoesNotQualify", err.Message)
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
	assert.Equal(t, "accountnumberDoesNotQualify", err.Message)
}

func TestGetUserDetailsRepository(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	userDetails, err := repo.GetUserDetails("1")
	assert.NotEmpty(t, userDetails)
	assert.Nil(t, err)
}

func TestGetUserDetailsErrorFetchingData(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	userDetails, err := repo.GetUserDetails("10")
	assert.Empty(t, userDetails)
	assert.NotNil(t, err)
	assert.Equal(t, "Error while fetching data", err.Message)
}

func TestUpdateUserDetailRepository(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	reqUpdate := DataRequestUpdateProfile{
		Nama:  "Binoto Manurung",
		NoHp:  "+66666666666",
		Email: "andreasjulyus@gmail.com",
		NoRek: "66666666",
	}

	updateUser := repo.UpdateUser("1", reqUpdate)
	assert.Empty(t, updateUser)
}

func TestUpdateUserError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := DataRequestUpdateProfile{
		Nama:  "Julyus Andreas",
		NoHp:  "+6281234567890",
		Email: "julyusmanurung@gmail.com",
		NoRek: "6181801052",
	}

	err := repo.UpdateUser("abc", req)
	assert.Equal(t, "strconv.ParseUint: parsing \"abc\": invalid syntax", err.Message)
	assert.NotNil(t, err)
}
