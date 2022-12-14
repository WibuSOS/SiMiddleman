package transaction

import (
	"net/http"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/product"
	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"
	"github.com/WibuSOS/sinarmas/backend/database"
	"github.com/WibuSOS/sinarmas/backend/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := database.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	return db
}

func newTestDB2(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{})
	assert.NoError(t, err)

	return db
}

func newTestDBError(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	return db
}

func TestUpdateStatusSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := RequestUpdateStatus{
		Status: "barang dibayar",
	}

	err := repo.UpdateStatusDelivery("1", req)
	assert.Empty(t, err)

}

func TestUpdateStatusErrorConvert(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	req := RequestUpdateStatus{
		Status: "barang dibayar",
	}

	err := repo.UpdateStatusDelivery("ase", req)
	assert.NotNil(t, err)

}

func TestGetDetailsPaymentSuccess(t *testing.T) {
	db := newTestDB2(t)
	repo := NewRepository(db)
	roomRepo := rooms.NewRepository(db)

	req := rooms.DataRequest{
		PenjualID: 1,
		Product: &product.DataRequest{
			Nama:      "Razer Mouse",
			Deskripsi: "Ini Razer Mouse",
			Harga:     150000,
			Kuantitas: 1,
		},
	}

	newRoom, err := roomRepo.CreateRoom(&req)
	assert.Empty(t, err)
	assert.NotEmpty(t, newRoom.RoomCode)

	paymentDetails, err := repo.GetPaymentDetails(int(newRoom.ID))
	assert.Nil(t, err)
	assert.NotNil(t, paymentDetails.Product)
}

func TestGetDetailsPaymentErrorFetchingData(t *testing.T) {
	db := newTestDBError(t)
	repo := NewRepository(db)

	_, err := repo.GetPaymentDetails(0)
	assert.NotNil(t, err)
	assert.Equal(t, "errorfetchingdata", err.Message)
	assert.Equal(t, http.StatusInternalServerError, err.Status)
	assert.Equal(t, "Internal_Server_Error", err.Error)
}

func TestGetDetailsRoomNotFound(t *testing.T) {
	db := newTestDB2(t)
	repo := NewRepository(db)

	idRoom := 3

	_, err := repo.GetPaymentDetails(idRoom)
	assert.NotNil(t, err)
	assert.Equal(t, "roomnotfound", err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}
