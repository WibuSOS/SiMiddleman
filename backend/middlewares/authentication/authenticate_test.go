package authentication

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/auth"
	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"
	"github.com/WibuSOS/sinarmas/backend/database"
	"github.com/WibuSOS/sinarmas/backend/middlewares/authorization"
	"github.com/WibuSOS/sinarmas/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type createRoomsResponse struct {
	Message string       `json:"message"`
	Data    models.Rooms `json:"data"`
}

func setEnv() {
	os.Setenv("ENVIRONMENT", "TEST")
}

func newTestDB(t *testing.T) *gorm.DB {
	db, err := database.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	p := "123456781234567812"
	hash, err := bcrypt.GenerateFromPassword([]byte(p), 10)
	assert.NoError(t, err)

	user := models.Users{Email: "consumer@xyz.com", Password: string(hash)}
	result := db.Create(&user)
	assert.NoError(t, result.Error)

	return db
}

func TestMain(m *testing.M) {
	setEnv()
	exitVal := m.Run()
	os.Exit(exitVal)
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

func newTestCreateRoomWithAuth(t *testing.T, withToken bool) (*httptest.ResponseRecorder, createRoomsResponse) {
	var createRoomsRes createRoomsResponse
	db := newTestDB(t)
	consumer := []string{"consumer"}
	consumerService := authorization.NewServiceAuthorization(db, consumer)
	consumerHandler := authorization.NewHandlerAuthorization(consumerService)

	// rooms Handler
	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	// Set Routes
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/rooms", Authentication, consumerHandler.RoleAuthorize, roomHandler.CreateRoom)

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
	if withToken {
		req.Header.Add("Authorization", token)
	}

	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w, createRoomsRes
}

func TestCreateRoomWithAuthSuccess(t *testing.T) {
	w, room := newTestCreateRoomWithAuth(t, true)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &room))
	assert.Equal(t, "success", room.Message)
}

func TestCreateRoomWithAuthNoToken(t *testing.T) {
	w, room := newTestCreateRoomWithAuth(t, false)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.NotEqual(t, "success", room.Message)
}
