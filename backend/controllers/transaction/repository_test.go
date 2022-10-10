package transaction

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/controllers/rooms"
	"github.com/WibuSOS/sinarmas/models"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{})
	assert.NoError(t, err)

	return db
}

func newTestCreateRoomWithProduct(t *testing.T) uint {
	type roomResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var resRoom roomResponse

	db := newTestDB(t)
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", roomHandler.CreateRoom)

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
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resRoom))
	assert.Equal(t, "success", resRoom.Message)
	assert.NotEmpty(t, resRoom.Data.RoomCode)

	return resRoom.Data.ID
}

func newTestCreateRoomWithoutProduct(t *testing.T) uint {
	type roomResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var resRoom roomResponse

	db := newTestDB(t)
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", roomHandler.CreateRoom)

	// SUCCESS
	payload := `{"id": 1}`
	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resRoom))
	assert.Equal(t, "success", resRoom.Message)
	assert.NotEmpty(t, resRoom.Data.RoomCode)

	return resRoom.Data.ID
}

func TestGetDetailsPaymentSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	idRoom := newTestCreateRoomWithProduct(t)

	paymentDetails, err := repo.GetPaymentDetails(int(idRoom))
	assert.Nil(t, err)
	assert.NotNil(t, paymentDetails.Product)
}

func TestGetDetailsProductEmpty(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	idRoom := newTestCreateRoomWithoutProduct(t)

	paymentDetails, err := repo.GetPaymentDetails(int(idRoom))
	assert.Nil(t, err)
	assert.Nil(t, paymentDetails.Product)
}

func TestGetDetailsRoomNotFound(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)

	idRoom := 3

	paymentDetails, err := repo.GetPaymentDetails(int(idRoom))
	assert.NotNil(t, err)
	assert.Nil(t, paymentDetails)
	assert.Equal(t, "Room Not Found", err.Message)
	assert.Equal(t, 400, err.Status)
	assert.Equal(t, "Bad_Request", err.Error)
}
