package transaction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStatusDelivery(t *testing.T) {

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	db.Create(&models.Products{
		RoomsID:   1,
		Nama:      "produk1",
		Harga:     15000,
		Kuantitas: 2,
		Deskripsi: "ini adalah produk 1",
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.PUT("/updatestatusdelivery/:id", handler.UpdateStatusDelivery)
	req, err := http.NewRequest("PUT", "/updatestatusdelivery/1", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type getUpdateStatusDelivery struct {
		Message string `json:"message"`
	}

	var res getUpdateStatusDelivery
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success update status pengiriman barang", res.Message)
}

func TestUpdateErrorStatusDelivery(t *testing.T) {

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	db.Create(&models.Products{
		RoomsID:   1,
		Nama:      "produk1",
		Harga:     15000,
		Kuantitas: 2,
		Deskripsi: "ini adalah produk 1",
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.PUT("/updatestatusdelivery/:id", handler.UpdateStatusDelivery)
	req, err := http.NewRequest("PUT", "/updatestatusdelivery/asasdaweq", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	type getUpdateStatusDelivery struct {
		Message string `json:"message"`
	}

	var res getUpdateStatusDelivery
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "WHERE conditions required", res.Message)
}
