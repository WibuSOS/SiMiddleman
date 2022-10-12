package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/controllers/rooms"
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
	payload := `{"status": "barang dibayar"}`
	req, err := http.NewRequest("PUT", "/updatestatusdelivery/1", strings.NewReader(payload))
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
	assert.Equal(t, "success update status", res.Message)
}

func TestUpdateStatusDeliveryInvalidJSON(t *testing.T) {

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
	payload := `{"status": 123}`
	req, err := http.NewRequest("PUT", "/updatestatusdelivery/1", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
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
	payload := `{"status": "barang dibayar"}`
	req, err := http.NewRequest("PUT", "/updatestatusdelivery/asasdaweq", strings.NewReader(payload))
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

func TestGetPaymentDetailsHandlerSuccess(t *testing.T) {
	db := newTestDB2(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", roomHandler.CreateRoom)
	r.GET("/getHarga/:idroom", handler.GetPaymentDetails)

	payload := `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`
	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type createRoomResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var createRoomRes createRoomResponse

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createRoomRes))
	assert.Equal(t, "success", createRoomRes.Message)
	assert.NotEmpty(t, createRoomRes.Data.RoomCode)

	url := fmt.Sprintf("/getHarga/%d", createRoomRes.Data.ID)
	req, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type response struct {
		Message string              `json:"message"`
		Data    ResponsePaymentInfo `json:"data"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Greater(t, int(res.Data.Total), 0)
}

func TestGetPaymentDetailsHandlerErrorQueryParam(t *testing.T) {
	db := newTestDB2(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getHarga/:idroom", handler.GetPaymentDetails)

	url := fmt.Sprintf("/getHarga/%s", "abc")
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPaymentDetailsHandlerRoomNotFound(t *testing.T) {
	db := newTestDB2(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getHarga/:idroom", handler.GetPaymentDetails)

	url := fmt.Sprintf("/getHarga/%d", 3)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
