package rooms

import (
	"encoding/json"
	"fmt"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/users"
	"github.com/WibuSOS/sinarmas/backend/models"
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

func TestJoinRoomPembeliHandlerFail(t *testing.T) {
	type response struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var res response

	db := newTestDB(t)
	// Rooms Handler
	roomsRepo := NewRepository(db)
	roomsService := NewService(roomsRepo)
	roomsHandler := NewHandler(roomsService)

	// Users Handler
	usersRepo := users.NewRepository(db)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register" /*authentication.Authentication, isAdmin.Authorize,*/, usersHandler.CreateUser)
	r.POST("/rooms", roomsHandler.CreateRoom)
	r.PUT("/en/joinroom/:room_id/:user_id" /*authentication.Authentication, isAdmin.Authorize,*/, roomsHandler.JoinRoomPembeli)

	// Create User 1
	payload := `{
		"nama": "klmno",
		"email": "admin@klm.com",
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
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)

	// Create User 2
	payload = `{
		"nama": "thea",
		"email": "admin@admin.com",
		"password": "123456781234567812",
		"noHp": "+6281223440777",
		"noRek": "1234"
	}`
	req, err = http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)

	// Create Room 1
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
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)

	// Fail join room not found pembeli
	req, err = http.NewRequest("PUT", "/en/joinroom/1/2", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.NotEmpty(t, res.Data)
}

func TestJoinRoomPembeliHandlerSuccess(t *testing.T) {
	type responseCreateRoom struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var resCreateRoom responseCreateRoom

	type responseJoinRoom struct {
		Message string `json:"message"`
	}
	var resJoinRoom responseJoinRoom

	db := newTestDB(t)
	// Rooms Handler
	roomsRepo := NewRepository(db)
	roomsService := NewService(roomsRepo)
	roomsHandler := NewHandler(roomsService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", roomsHandler.CreateRoom)
	r.PUT("/en/joinroom/:room_id/:user_id" /*authentication.Authentication, isAdmin.Authorize,*/, roomsHandler.JoinRoomPembeli)

	// Create Room 1
	payload := (`{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`)

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resCreateRoom))
	assert.Equal(t, "success", resCreateRoom.Message)

	roomCode := resCreateRoom.Data.RoomCode

	url := fmt.Sprintf("/en/joinroom/%s/%d", roomCode, 2)
	req, err = http.NewRequest("PUT", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resJoinRoom))
	assert.Equal(t, "success", resJoinRoom.Message)
}

func TestJoinRoomPembeliHandlerAlreadyInRoom(t *testing.T) {
	type responseCreateRoom struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var resCreateRoom responseCreateRoom

	type responseJoinRoom struct {
		Message string `json:"message"`
	}
	var resJoinRoom responseJoinRoom

	db := newTestDB(t)
	// Rooms Handler
	roomsRepo := NewRepository(db)
	roomsService := NewService(roomsRepo)
	roomsHandler := NewHandler(roomsService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", roomsHandler.CreateRoom)
	r.PUT("/en/joinroom/:room_id/:user_id" /*authentication.Authentication, isAdmin.Authorize,*/, roomsHandler.JoinRoomPembeli)

	// Create Room 1
	payload := (`{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`)

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resCreateRoom))
	assert.Equal(t, "success", resCreateRoom.Message)

	roomCode := resCreateRoom.Data.RoomCode

	url := fmt.Sprintf("/en/joinroom/%s/%d", roomCode, 1)
	req, err = http.NewRequest("PUT", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resJoinRoom))
}

func TestJoinRoomPembeliHandlerInvalidUserID(t *testing.T) {
	type responseCreateRoom struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var resCreateRoom responseCreateRoom

	type responseJoinRoom struct {
		Message string `json:"message"`
	}
	var resJoinRoom responseJoinRoom

	db := newTestDB(t)
	// Rooms Handler
	roomsRepo := NewRepository(db)
	roomsService := NewService(roomsRepo)
	roomsHandler := NewHandler(roomsService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", roomsHandler.CreateRoom)
	r.PUT("/en/joinroom/:room_id/:user_id" /*authentication.Authentication, isAdmin.Authorize,*/, roomsHandler.JoinRoomPembeli)

	// Create Room 1
	payload := (`{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`)

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resCreateRoom))
	assert.Equal(t, "success", resCreateRoom.Message)

	roomCode := resCreateRoom.Data.RoomCode

	url := fmt.Sprintf("/en/joinroom/%s/%s", roomCode, "abc")
	req, err = http.NewRequest("PUT", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resJoinRoom))
}

func TestJoinRoomSuccess(t *testing.T) {
	type responseCreateRoom struct {
		Message string       `json:"message"`
		Data    models.Rooms `json:"data"`
	}
	var resCreateRoom responseCreateRoom

	type responseJoinRoom struct {
		Message string `json:"message"`
	}
	var resJoinRoom responseJoinRoom

	db := newTestDB(t)
	// Rooms Handler
	roomsRepo := NewRepository(db)
	roomsService := NewService(roomsRepo)
	roomsHandler := NewHandler(roomsService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", roomsHandler.CreateRoom)
	r.GET("/en/joinroom/:room_id/:user_id" /*authentication.Authentication, isAdmin.Authorize,*/, roomsHandler.JoinRoom)

	// Create Room 1
	payload := (`{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`)

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resCreateRoom))
	assert.Equal(t, "success", resCreateRoom.Message)

	url := fmt.Sprintf("/en/joinroom/%v/%v", "1", "1")
	req, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &resJoinRoom))
}
