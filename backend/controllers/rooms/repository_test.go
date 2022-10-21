package rooms

import (
	"net/http"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/product"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{})
	assert.NoError(t, err)

	res := db.Exec("PRAGMA foreign_keys = ON", nil)
	assert.NoError(t, res.Error)

	// USER 1
	pb, _ := bcrypt.GenerateFromPassword([]byte("123456781234567812"), 8)
	newUser := models.Users{
		Nama:     "vwxyz",
		Role:     "consumer",
		NoHp:     "+6283785332789",
		Email:    "admin@xyz.com",
		Password: string(pb),
		NoRek:    "1234",
	}
	res = db.Omit(clause.Associations).Create(&newUser)
	assert.NoError(t, res.Error)

	// USER 2
	pb, _ = bcrypt.GenerateFromPassword([]byte("123456781234567812"), 8)
	newUser = models.Users{
		Nama:     "abcde",
		Role:     "consumer",
		NoHp:     "+6282876443890",
		Email:    "admin@abc.com",
		Password: string(pb),
		NoRek:    "6789",
	}
	res = db.Omit(clause.Associations).Create(&newUser)
	assert.NoError(t, res.Error)

	// USER 3
	pb, _ = bcrypt.GenerateFromPassword([]byte("123456781234567812"), 8)
	newUser = models.Users{
		Nama:     "fghij",
		Role:     "consumer",
		NoHp:     "+6283987554901",
		Email:    "admin@fgh.com",
		Password: string(pb),
		NoRek:    "5678",
	}
	res = db.Omit(clause.Associations).Create(&newUser)
	assert.NoError(t, res.Error)

	return db
}

func TestCreateRoomRepositorySuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// SUCCESS
	req := DataRequest{
		PenjualID: 1,
		Product: &product.DataRequest{
			Nama:      "Razer Mouse",
			Deskripsi: "Ini Razer Mouse",
			Harga:     150000,
			Kuantitas: 1,
		},
	}

	newRoom, err := repo.CreateRoom(&req)
	assert.Empty(t, err)
	assert.NotEmpty(t, newRoom.RoomCode)
}

func TestCreateRoomRepositoryErrorDbLevel(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ERROR PENJUAL ID
	req := DataRequest{
		PenjualID: 10,
		Product: &product.DataRequest{
			Nama:      "Razer Mouse",
			Deskripsi: "Ini Razer Mouse",
			Harga:     150000,
			Kuantitas: 1,
		},
	}

	newRoom, err := repo.CreateRoom(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "badRequest", err.Message)
	assert.Empty(t, newRoom.RoomCode)
}

func TestCreateRoomRepositoryErrorProduct(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ERROR NAMA
	req := DataRequest{
		PenjualID: 1,
		Product: &product.DataRequest{
			Nama:      "",
			Deskripsi: "Ini Razer Mouse",
			Harga:     150000,
			Kuantitas: 1,
		},
	}

	newRoom, err := repo.CreateRoom(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "emptyname", err.Message)
	assert.Empty(t, newRoom.RoomCode)
}

func TestGetAllRoomsRepositorySuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	// ROOM 1
	req := DataRequest{
		PenjualID: 1,
		Product: &product.DataRequest{
			Nama:      "Razer Mouse",
			Deskripsi: "Ini Razer Mouse",
			Harga:     150000,
			Kuantitas: 1,
		},
	}

	newRoom, err := repo.CreateRoom(&req)
	assert.Empty(t, err)
	assert.NotEmpty(t, newRoom.RoomCode)

	// ROOM 2
	req = DataRequest{
		PenjualID: 1,
		Product: &product.DataRequest{
			Nama:      "Razer Mouse",
			Deskripsi: "Ini Razer Mouse",
			Harga:     150000,
			Kuantitas: 1,
		},
	}

	newRoom, err = repo.CreateRoom(&req)
	assert.Empty(t, err)
	assert.NotEmpty(t, newRoom.RoomCode)

	rooms, err := repo.GetAllRooms("1")
	assert.Empty(t, err)
	assert.NotEmpty(t, rooms)
	assert.Len(t, rooms, 2)
}

func TestGetAllRoomsRepositoryErrorDbLevel(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	rooms, err := repo.GetAllRooms("10")
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Empty(t, rooms)
}
