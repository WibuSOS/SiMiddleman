package authorization

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/controllers/auth"
	"github.com/WibuSOS/sinarmas/controllers/rooms"
	"github.com/WibuSOS/sinarmas/controllers/users"
	"github.com/WibuSOS/sinarmas/middlewares/authentication"
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

func TestCreateRoomWithAuthSuccess(t *testing.T) {
	type createUsersResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	type loginResponse struct {
		Message string            `json:"message"`
		Data    auth.DataResponse `json:"data"`
		Token   string            `json:"token"`
	}
	type createRoomsResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}

	var createUsersRes createUsersResponse
	var createRoomsRes createRoomsResponse
	var loginRes loginResponse
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	db := newTestDB(t)
	consumer := []string{"consumer"}
	isConsumer := Roles{AllowedRoles: consumer[:]}

	// Users Handler
	usersRepo := users.NewRepository(db)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	// Auth Handler
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// rooms Handler
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", usersHandler.CreateUser)
	r.POST("/login", authHandler.Login)
	r.POST("/rooms", authentication.Authentication, isConsumer.Authorize, roomHandler.CreateRoom)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// Create User
	payload := `{
		"nama": "xyzde",
		"email": "admin@hij.com",
		"password": "123456781234567812",
		"noHp": "+6281223440777",
		"noRek": "1234"
			}`

	req, err := http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createUsersRes))
	assert.Equal(t, "success", createUsersRes.Message)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// Login
	payload = `{"email": "admin@hij.com", "password": "123456781234567812"}`
	req, err = http.NewRequest("POST", "/login", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &loginRes))
	assert.Equal(t, "success", loginRes.Message)
	assert.NotEqual(t, "", loginRes.Data.Email)
	assert.NotEqual(t, "", loginRes.Token)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// SUCCESS
	payload = `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	req, err = http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", loginRes.Token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createRoomsRes))
	assert.Equal(t, "success", createRoomsRes.Message)
}

func TestCreateRoomWithAuthUnauthorize(t *testing.T) {
	type createUsersResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	type loginResponse struct {
		Message string            `json:"message"`
		Data    auth.DataResponse `json:"data"`
		Token   string            `json:"token"`
	}
	type createRoomsResponse struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}

	var createUsersRes createUsersResponse
	var createRoomsRes createRoomsResponse
	var loginRes loginResponse
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	db := newTestDB(t)
	admin := []string{"admin"}
	isAdmin := Roles{AllowedRoles: admin[:]}

	// Users Handler
	usersRepo := users.NewRepository(db)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	// Auth Handler
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// rooms Handler
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", usersHandler.CreateUser)
	r.POST("/login", authHandler.Login)
	r.POST("/rooms", authentication.Authentication, isAdmin.Authorize, roomHandler.CreateRoom)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// Create User
	payload := `{
		"nama": "xyzde",
		"email": "admin@pqr.com",
		"password": "123456781234567812",
		"noHp": "+6281223440777",
		"noRek": "1234"
			}`

	req, err := http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createUsersRes))
	assert.Equal(t, "success", createUsersRes.Message)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// Login
	payload = `{"email": "admin@pqr.com", "password": "123456781234567812"}`
	req, err = http.NewRequest("POST", "/login", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	log.Println(&loginRes)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &loginRes))
	assert.Equal(t, "success", loginRes.Message)
	assert.NotEqual(t, "", loginRes.Data.Email)
	assert.NotEqual(t, "", loginRes.Token)
	//--------------------------------------------------------------------------------------------------------------------------------------------//
	// SUCCESS
	payload = `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	req, err = http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", loginRes.Token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEqual(t, "success", createRoomsRes.Message)
}
