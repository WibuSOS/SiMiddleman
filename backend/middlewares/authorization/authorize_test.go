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

var (
	consumer  = []string{"consumer"}
	admin     = []string{"admin"}
	consumer1 = "consumer@xyz.com"
	consumer2 = "consumer@abc.com"
	consumer3 = "consumer@ijk.com"
	p         = "123456781234567812"
)

type loginResponse struct {
	Message string            `json:"message"`
	Data    auth.DataResponse `json:"data"`
	Token   string            `json:"token"`
}

type createRoomResponse struct {
	Message string       `json:"message"`
	Data    models.Rooms `json:"data"`
}

type getPaymentDetailsResponse struct {
	Message string                          `json:"message"`
	Data    transaction.ResponsePaymentInfo `json:"data"`
}

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{}, &models.Rooms{}, &models.Products{}, &models.Transactions{})
	assert.NoError(t, err)

	consumerArr := []string{consumer1, consumer2, consumer3}
	for i := 0; i < len(consumerArr); i++ {
		hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
		assert.NoError(t, err)

		user := models.Users{Email: consumerArr[i], Password: string(hash)}
		result := db.Create(&user)
		assert.NoError(t, result.Error)
	}

	return db
}

func newTestLoginHandler(t *testing.T, email string) loginResponse {
	var res loginResponse

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

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.NotEqual(t, "", res.Data.Email)
	assert.NotEqual(t, "", res.Token)

	return res
}

func newTestGetPaymentDetailsHandlerWithRoomAuth(t *testing.T, allowedRoles []string, withID bool, roomID *string) (*httptest.ResponseRecorder, getPaymentDetailsResponse) {
	var createRoomRes createRoomResponse
	var res getPaymentDetailsResponse

	db := newTestDB(t)

	reqService := NewServiceAuthorization(db, allowedRoles)
	reqHandler := NewHandlerAuthorization(reqService)

	consumerService := NewServiceAuthorization(db, consumer)
	consumerHandler := NewHandlerAuthorization(consumerService)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	handler := transaction.NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", authentication.Authentication, consumerHandler.RoleAuthorize, roomHandler.CreateRoom)
	r.GET("/getHarga/:room_id", authentication.Authentication, reqHandler.RoleAuthorize, reqHandler.RoomAuthorize, handler.GetPaymentDetails)

	payload := `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	login := newTestLoginHandler(t, consumer1)

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", login.Token)
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createRoomRes))
	assert.Equal(t, "success", createRoomRes.Message)
	assert.NotEmpty(t, createRoomRes.Data.RoomCode)

	var url string
	if withID {
		url = fmt.Sprintf("/getHarga/%d", createRoomRes.Data.ID)
	} else {
		if roomID != nil {
			url = fmt.Sprintf("/getHarga/%s", *roomID)
		} else {
			url = fmt.Sprintf("/getHarga/%s", createRoomRes.Data.RoomCode)
		}
	}

	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", login.Token)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w, res
}

func TestGetPaymentDetailsHandlerWithRoomAuthSuccess(t *testing.T) {
	w, paymentDetails := newTestGetPaymentDetailsHandlerWithRoomAuth(t, consumer, true, nil)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &paymentDetails))
	assert.Equal(t, "success", paymentDetails.Message)
	assert.Greater(t, int(paymentDetails.Data.Total), 0)
}

func TestGetPaymentDetailsHandlerWithRoomAuthInvalidRoomID(t *testing.T) {
	w, paymentDetails := newTestGetPaymentDetailsHandlerWithRoomAuth(t, consumer, false, nil)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &paymentDetails))
	assert.Equal(t, "Invalid Room ID", paymentDetails.Message)
}

func TestGetPaymentDetailsHandlerWithRoomAuthRoomNotFound(t *testing.T) {
	roomID := "7"
	w, paymentDetails := newTestGetPaymentDetailsHandlerWithRoomAuth(t, consumer, false, &roomID)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &paymentDetails))
	assert.Equal(t, "Room not found", paymentDetails.Message)
}

func TestGetPaymentDetailsHandlerWithRoomAuthUnauthorizeRole(t *testing.T) {
	w, paymentDetails := newTestGetPaymentDetailsHandlerWithRoomAuth(t, admin, true, nil)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &paymentDetails))
	assert.Equal(t, "Unauthorized", paymentDetails.Message)
}

func TestGetPaymentDetailsHandlerWithRoomAuthUnauthorizeRoom(t *testing.T) {
	w, paymentDetails := newTestGetPaymentDetailsHandlerWithRoomAuth(t, admin, true, nil)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &paymentDetails))
	assert.Equal(t, "Unauthorized", paymentDetails.Message)
}

func TestGetPaymentDetailsHandlerWithRoomAuthUnauthorize(t *testing.T) {
	var createRoomRes createRoomResponse
	db := newTestDB(t)

	consumerService := NewServiceAuthorization(db, consumer)
	consumerHandler := NewHandlerAuthorization(consumerService)

	repo := transaction.NewRepository(db)
	service := transaction.NewService(repo)
	handler := transaction.NewHandler(service)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", authentication.Authentication, consumerHandler.RoleAuthorize, roomHandler.CreateRoom)
	r.PUT("/joinroom/:room_id/:user_id", authentication.Authentication, consumerHandler.RoleAuthorize, roomHandler.JoinRoomPembeli)
	r.GET("/getHarga/:room_id", authentication.Authentication, consumerHandler.RoleAuthorize, consumerHandler.RoomAuthorize, handler.GetPaymentDetails)

	payload := `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	login := newTestLoginHandler(t, consumer1)

	req, err := http.NewRequest("POST", "/rooms", strings.NewReader(payload))
	req.Header.Add("Authorization", login.Token)
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createRoomRes))
	assert.Equal(t, "success", createRoomRes.Message)
	assert.NotEmpty(t, createRoomRes.Data.RoomCode)

	login = newTestLoginHandler(t, consumer2)

	roomCode := createRoomRes.Data.RoomCode

	url := fmt.Sprintf("/joinroom/%s/%d", roomCode, 2)
	req, err = http.NewRequest("PUT", url, nil)
	req.Header.Add("Authorization", login.Token)
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

	login = newTestLoginHandler(t, consumer3)

	url = fmt.Sprintf("/getHarga/%d", createRoomRes.Data.ID)
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", login.Token)
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
