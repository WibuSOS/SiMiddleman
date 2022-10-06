package rooms

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

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

	err = db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{})
	assert.NoError(t, err)

	return db
}

func TestCreateRoom(t *testing.T) {
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
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)

	// ERROR CREATE ROOM
	payload = `{
		"id": 0,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`
	req, err = http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "oops... there is something wrong", res.Message)

	// ERROR BIND
	payload = `{}`
	req, err = http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// ERROR VALIDATE USER ID
	payload = `{
		"id": "1",
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`
	req, err = http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	// ERROR VALIDATE PRODUCTS
	payload = `{
		"id": 1,
		"product" : {
			"nama": "",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`
	req, err = http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
