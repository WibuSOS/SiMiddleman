package authentication

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/controllers/auth"
	"github.com/WibuSOS/sinarmas/controllers/rooms"
	"github.com/WibuSOS/sinarmas/middlewares/authorization"
	"github.com/WibuSOS/sinarmas/models"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{})
	assert.NoError(t, err)

	password := "123456781234567812"
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	assert.NoError(t, err)

	user := models.Users{Email: "consumer@xyz.com", Password: string(hash)}
	result := db.Create(&user)
	assert.NoError(t, result.Error)

	return db
}

func newTestLoginHandler(t *testing.T) string {
	db := newTestDB(t)
	repo := auth.NewRepository(db)
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/login", handler.Login)
	payload := `{"email": "consumer@xyz.com", "password": "123456781234567812"}`
	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type response struct {
		Message string            `json:"message"`
		Data    auth.DataResponse `json:"data"`
		Token   string            `json:"token"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.NotEqual(t, "", res.Data.Email)
	assert.NotEqual(t, "", res.Token)

	return res.Token
}

func TestCreateRoomWithAuthSuccess(t *testing.T) {
	type createRoomsResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}

	var createRoomsRes createRoomsResponse
	db := newTestDB(t)
	consumer := []string{"consumer"}
	isConsumer := authorization.Roles{AllowedRoles: consumer[:]}

	// rooms Handler
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)
	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", Authentication, isConsumer.Authorize, roomHandler.CreateRoom)
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

	token := newTestLoginHandler(t)

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createRoomsRes))
	assert.Equal(t, "success", createRoomsRes.Message)
}

func TestCreateRoomWithAuthNoToken(t *testing.T) {
	type createRoomsResponse struct {
		Message string `json:"message"`
	}

	var createRoomsRes createRoomsResponse
	db := newTestDB(t)
	consumer := []string{"consumer"}
	isConsumer := authorization.Roles{AllowedRoles: consumer[:]}

	// rooms Handler
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)
	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", Authentication, isConsumer.Authorize, roomHandler.CreateRoom)
	// No Token
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

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEqual(t, "success", createRoomsRes.Message)
}
