package rooms

import (
	"net/http"
	"testing"

	"github.com/WibuSOS/sinarmas/controllers/product"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoomServiceSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

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

	newRoom, err := service.CreateRoom(&req)
	assert.Empty(t, err)
	assert.NotEmpty(t, newRoom.RoomCode)
}

func TestCreateRoomServiceError(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)

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

	newRoom, err := service.CreateRoom(&req)
	assert.NotEmpty(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Equal(t, "constraint failed: FOREIGN KEY constraint failed (787)", err.Message)
	assert.Empty(t, newRoom.RoomCode)
}
