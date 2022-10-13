package rooms

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/users"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestJoinRoom(t *testing.T) {
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
	r.GET("/joinroom/:room_id/:user_id" /*authentication.Authentication, isAdmin.Authorize,*/, roomsHandler.JoinRoom)

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

	// Success join room
	req, err = http.NewRequest("GET", "/joinroom/1/1", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.NotEmpty(t, res.Data)

	// Id penjual / pembeli tidak sesuai
	req, err = http.NewRequest("GET", "/joinroom/1/2", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Tidak bisa memasuki ruangan", res.Message)
	assert.NotEmpty(t, res.Data)

	// Error Inputan salah
	req, err = http.NewRequest("GET", "/joinroom/~@/test123", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Tidak bisa memasuki ruangan", res.Message)
	assert.NotEmpty(t, res.Data)

}
