package transaction

import (
	"net/http"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/product"
	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"

	"github.com/stretchr/testify/assert"
)

func TestServiceUpdateStatus(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := RequestUpdateStatus{
		Status: "barang dibayar",
	}

	err := service.UpdateStatusDelivery("1", req)
	assert.Nil(t, err)

}

func TestServiceUpdateStatusInvalidStatus(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := RequestUpdateStatus{
		Status: "test",
	}

	err := service.UpdateStatusDelivery("1", req)
	assert.NotNil(t, err)
	assert.Equal(t, "invalidstatus", err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

func TestServiceUpdateStatusError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

	req := RequestUpdateStatus{
		Status: "barang dibayar",
	}

	err := service.UpdateStatusDelivery("asd", req)
	assert.NotNil(t, err)
	assert.Equal(t, "WHERE conditions required", err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}

func TestGetPaymentDetailsServiceSuccess(t *testing.T) {
	db := newTestDB2(t)
	repo := NewRepository(db)
	service := NewService(repo)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)

	req := rooms.DataRequest{
		PenjualID: 1,
		Product: &product.DataRequest{
			Nama:      "Razer Mouse",
			Deskripsi: "Ini Razer Mouse",
			Harga:     150000,
			Kuantitas: 1,
		},
	}

	newRoom, err := roomService.CreateRoom(&req)
	assert.Empty(t, err)
	assert.NotEmpty(t, newRoom.RoomCode)

	res, err := service.GetPaymentDetails(int(newRoom.ID))
	assert.Nil(t, err)
	assert.Greater(t, int(res.Total), 0)
}

func TestGetPaymentDetailsServiceRoomNotFound(t *testing.T) {
	db := newTestDB2(t)
	repo := NewRepository(db)
	service := NewService(repo)

	idRoom := 3

	_, err := service.GetPaymentDetails(int(idRoom))
	assert.NotNil(t, err)
	assert.Equal(t, "roomnotfound", err.Message)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}
