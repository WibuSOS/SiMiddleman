package authorization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/auth"
	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"
	"github.com/WibuSOS/sinarmas/backend/controllers/transaction"
	"github.com/WibuSOS/sinarmas/backend/middlewares/authentication"
	"github.com/WibuSOS/sinarmas/backend/models"

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

	// User 1
	p := "123456781234567812"
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	assert.NoError(t, err)

	user := models.Users{Email: "consumer@xyz.com", Password: string(hash)}
	result := db.Create(&user)
	assert.NoError(t, result.Error)

	// User 2
	p = "123456781234567812"
	hash, err = bcrypt.GenerateFromPassword([]byte(p), 10)
	assert.NoError(t, err)

	user = models.Users{Email: "consumer@abc.com", Password: string(hash)}
	result = db.Create(&user)
	assert.NoError(t, result.Error)

	// User 3
	p = "123456781234567812"
	hash, err = bcrypt.GenerateFromPassword([]byte(p), 10)
	assert.NoError(t, err)

	user = models.Users{Email: "consumer@ijk.com", Password: string(hash)}
	result = db.Create(&user)
	assert.NoError(t, result.Error)

	return db
}

func newTestLoginHandler(t *testing.T, email string) string {
	db := newTestDB(t)
	repo := auth.NewRepository(db)
	service := auth.NewService(repo)
	handler := auth.NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/login", handler.Login)
	payload := fmt.Sprintf(`{"email": "%s", "password": "123456781234567812"}`, email)
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
	isConsumer := Roles{AllowedRoles: consumer[:]}

	// rooms Handler
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", authentication.Authentication, isConsumer.Authorize, roomHandler.CreateRoom)

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

	token := newTestLoginHandler(t, "consumer@xyz.com")

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

func TestCreateRoomWithAuthUnauthorize(t *testing.T) {
	type createRoomsResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}

	var createRoomsRes createRoomsResponse

	db := newTestDB(t)
	admin := []string{"admin"}
	isAdmin := Roles{AllowedRoles: admin[:]}

	// rooms Handler
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", authentication.Authentication, isAdmin.Authorize, roomHandler.CreateRoom)

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

	token := newTestLoginHandler(t, "consumer@xyz.com")

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEqual(t, "success", createRoomsRes.Message)
}

func TestGetPaymentDetailsHandlerWithRoomAuthSuccess(t *testing.T) {
	db := newTestDB(t)

	roomAuth := NewRoomAuth(db)
	consumer := []string{"consumer"}
	isConsumer := Roles{AllowedRoles: consumer[:]}

	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	handler := transaction.NewHandler(service)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", authentication.Authentication, isConsumer.Authorize, roomHandler.CreateRoom)
	r.GET("/getHarga/:room_id", authentication.Authentication, isConsumer.Authorize, roomAuth.RoomAuthorize, handler.GetPaymentDetails)

	payload := `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	token := newTestLoginHandler(t, "consumer@xyz.com")

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", token)
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
	req.Header.Add("Authorization", token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type response struct {
		Message string                          `json:"message"`
		Data    transaction.ResponsePaymentInfo `json:"data"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Greater(t, int(res.Data.Total), 0)
}

func TestGetPaymentDetailsHandlerWithRoomAuthUnauthorize(t *testing.T) {
	db := newTestDB(t)

	roomAuth := NewRoomAuth(db)
	consumer := []string{"consumer"}
	isConsumer := Roles{AllowedRoles: consumer[:]}

	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	handler := transaction.NewHandler(service)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", authentication.Authentication, isConsumer.Authorize, roomHandler.CreateRoom)
	r.PUT("/joinroom/:room_id/:user_id", authentication.Authentication, isConsumer.Authorize, roomHandler.JoinRoomPembeli)
	r.GET("/getHarga/:room_id", authentication.Authentication, isConsumer.Authorize, roomAuth.RoomAuthorize, handler.GetPaymentDetails)

	payload := `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	token := newTestLoginHandler(t, "consumer@xyz.com")

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", token)
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

	token = newTestLoginHandler(t, "consumer@abc.com")

	roomCode := createRoomRes.Data.RoomCode

	url := fmt.Sprintf("/joinroom/%s/%d", roomCode, 2)
	req, err = http.NewRequest("PUT", url, nil)
	req.Header.Add("Authorization", token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type joinRoomResponse struct {
		Message string `json:"message"`
	}
	var joinRoomRes joinRoomResponse

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &joinRoomRes))
	assert.Equal(t, "success", joinRoomRes.Message)

	token = newTestLoginHandler(t, "consumer@ijk.com")

	url = fmt.Sprintf("/getHarga/%d", createRoomRes.Data.ID)
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	type response struct {
		Message string `json:"message"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Unauthorized", res.Message)
}

func TestGetPaymentDetailsHandlerWithRoomAuthInvalidRoomID(t *testing.T) {
	db := newTestDB(t)

	roomAuth := NewRoomAuth(db)
	consumer := []string{"consumer"}
	isConsumer := Roles{AllowedRoles: consumer[:]}

	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	handler := transaction.NewHandler(service)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", authentication.Authentication, isConsumer.Authorize, roomHandler.CreateRoom)
	r.GET("/getHarga/:room_id", authentication.Authentication, isConsumer.Authorize, roomAuth.RoomAuthorize, handler.GetPaymentDetails)

	payload := `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	token := newTestLoginHandler(t, "consumer@xyz.com")

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", token)
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

	url := fmt.Sprintf("/getHarga/%s", createRoomRes.Data.RoomCode)
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	type response struct {
		Message string `json:"message"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Invalid Room ID", res.Message)
}

func TestGetPaymentDetailsHandlerWithRoomAuthRoomNotFound(t *testing.T) {
	db := newTestDB(t)

	roomAuth := NewRoomAuth(db)
	consumer := []string{"consumer"}
	isConsumer := Roles{AllowedRoles: consumer[:]}

	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	handler := transaction.NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getHarga/:room_id", authentication.Authentication, isConsumer.Authorize, roomAuth.RoomAuthorize, handler.GetPaymentDetails)

	token := newTestLoginHandler(t, "consumer@xyz.com")

	url := fmt.Sprintf("/getHarga/%d", 7)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	type response struct {
		Message string `json:"message"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Room not found", res.Message)
}
