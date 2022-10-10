package rooms

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateRoomHandlerSuccess(t *testing.T) {
	type response struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var res response

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", handler.CreateRoom)

	// SUCCESS
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
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.NotEmpty(t, res.Data.RoomCode)
}

func TestCreateRoomHandlerErrorBind(t *testing.T) {
	type response struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var res response

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", handler.CreateRoom)

	// ERROR BIND
	payload := `{
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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Key: 'DataRequest.PenjualID' Error:Field validation for 'PenjualID' failed on the 'required' tag", res.Message)
	assert.Empty(t, res.Data.RoomCode)
}

func TestCreateRoomHandlerErrorRequest(t *testing.T) {
	type response struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var res response

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", handler.CreateRoom)

	// ERROR PENJUAL ID
	payload := `{
		"id": 10,
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

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "constraint failed: FOREIGN KEY constraint failed (787)", res.Message)
	assert.Empty(t, res.Data.RoomCode)
}

func TestGetAllRoomsHandlerSuccess(t *testing.T) {
	type getAllRoomsResponse struct {
		Message string         `json:"message"`
		Data    []models.Rooms `json:"data"`
	}
	type createRoomResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var getAllRoomsRes getAllRoomsResponse
	var createRoomRes createRoomResponse

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", handler.CreateRoom)
	r.GET("/rooms/:id", handler.GetAllRooms)

	// ROOM 1
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
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createRoomRes))
	assert.Equal(t, "success", createRoomRes.Message)
	assert.NotEmpty(t, createRoomRes.Data.RoomCode)

	// SUCCESS ADA ISINYA
	req, err = http.NewRequest("GET", "/rooms/1", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &getAllRoomsRes))
	assert.Equal(t, "success", getAllRoomsRes.Message)
	assert.NotEmpty(t, getAllRoomsRes.Data)
	assert.Len(t, getAllRoomsRes.Data, 1)
}

func TestGetAllRoomsHandlerErrorRequest(t *testing.T) {
	type getAllRoomsResponse struct {
		Message string         `json:"message"`
		Data    []models.Rooms `json:"data"`
	}
	var getAllRoomsRes getAllRoomsResponse

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/rooms/:id", handler.GetAllRooms)

	// ERROR RECORD NOT FOUND
	req, err := http.NewRequest("GET", "/rooms/10", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &getAllRoomsRes))
	assert.Equal(t, "record not found", getAllRoomsRes.Message)
	assert.Empty(t, getAllRoomsRes.Data)
}