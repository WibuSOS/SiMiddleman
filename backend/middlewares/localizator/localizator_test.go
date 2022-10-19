package localizator

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/api"
	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type response struct {
	Message string       `json:"message"`
	Data    models.Rooms `json:"data"`
}

// func setLog(t *testing.T) {
// 	file, err := os.OpenFile("./logs_test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
// 	assert.NoError(t, err)
// 	log.SetOutput(file)
// }

func newTestDB(t *testing.T) *gorm.DB {
	db, err := api.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	return db
}

func setEndPointHandler(t *testing.T, db *gorm.DB) *rooms.Handler {
	repo := rooms.NewRepository(db)
	assert.NotNil(t, repo)
	service := rooms.NewService(repo)
	assert.NotNil(t, service)
	handler := rooms.NewHandler(service)
	assert.NotNil(t, handler)

	return handler
}

func setLocalizationHandler(t *testing.T) *Handler {
	handler, err := NewHandler()
	assert.NotNil(t, handler)
	assert.NoError(t, err)

	return handler
}

func TestLocalizeSuccess(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")

	// DB INITIALIZATION
	db := newTestDB(t)

	var res response

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(localizationHandler.PassLocalizator)
	r.GET("/:lang/joinroom/:room_id/:user_id", endPointHandler.JoinRoom)

	// SUCCESS
	req, err := http.NewRequest("GET", "/id/joinroom/1/1", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "sukses", res.Message)
}

// func TestLocalizeFail(t *testing.T) {
// 	os.Setenv("ENVIRONMENT", "TEST")

// 	// DB INITIALIZATION
// 	db := newTestDB(t)

// 	var res response

// 	// LOCALIZATION HANDLER
// 	localizationHandler := setLocalizationHandler(t)

// 	// END-POINT HANDLER
// 	endPointHandler := setEndPointHandler(t, db)

// 	// ROUTES INITIALIZATION
// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default()
// 	r.Use(localizationHandler.PassLocalizator)
// 	r.GET("/:lang/joinroom/:room_id/:user_id", endPointHandler.JoinRoom)

// 	// SUCCESS
// 	req, err := http.NewRequest("GET", "/id/joinroom/1/1", nil)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, req)

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
// 	assert.Equal(t, "sukses", res.Message)
// }