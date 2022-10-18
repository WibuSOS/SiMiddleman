package localizator

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/api"
	"github.com/WibuSOS/sinarmas/backend/controllers/users"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type response struct {
	Message string `json:"message"`
}

func TestLocalizeSuccess(t *testing.T) {
	// DB INITIALIZATION
	os.Setenv("ENVIRONMENT", "TEST")
	db, err := api.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	var res response

	// LOCALIZATION HANDLER
	localizatorHandler := NewHandler()
	assert.NotNil(t, localizatorHandler)

	// USERS HANDLER
	usersRepo := users.NewRepository(db)
	assert.NotNil(t, usersRepo)
	usersService := users.NewService(usersRepo)
	assert.NotNil(t, usersService)
	usersHandler := users.NewHandler(usersService)
	assert.NotNil(t, usersHandler)

	// ROUTES INITIALIZATION
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", localizatorHandler.PassLocalizator, usersHandler.CreateUser)

	// SUCCESS
	payload := `{
		"nama":     "vwxyz",
		"noHp":     "+6283785332789",
		"email":    "admin@xyz.com",
		"password": "123456781234567812",
		"noRek":    "1234"
	}`
	req, err := http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
}

func TestLocalizeFail(t *testing.T) {
	// DB INITIALIZATION
	os.Setenv("ENVIRONMENT", "TEST")
	db, err := api.SetupDb()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	var res response

	// LOCALIZATION HANDLER
	localizatorHandler := NewHandler()
	assert.NotNil(t, localizatorHandler)

	// USERS HANDLER
	usersRepo := users.NewRepository(db)
	assert.NotNil(t, usersRepo)
	usersService := users.NewService(usersRepo)
	assert.NotNil(t, usersService)
	usersHandler := users.NewHandler(usersService)
	assert.NotNil(t, usersHandler)

	// ROUTES INITIALIZATION
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", localizatorHandler.PassLocalizator, usersHandler.CreateUser)

	// SUCCESS
	payload := `{
		"nama":     "vwxyz",
		"noHp":     "+6283785332789",
		"email":    "admin@xyz.com",
		"password": "123456781234567812",
		"noRek":    "1234"
	}`
	req, err := http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
}
