package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/middlewares/localizator"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type response struct {
	Message string `json:"message"`
}

func setEndPointHandler(t *testing.T, db *gorm.DB) *Handler {
	repo := NewRepository(db)
	assert.NotNil(t, repo)
	service := NewService(repo)
	assert.NotNil(t, service)
	handler := NewHandler(service)
	assert.NotNil(t, handler)

	return handler
}

func setLocalizationHandler(t *testing.T) *localizator.Handler {
	handler, err := localizator.NewHandler()
	assert.NotNil(t, handler)
	assert.NoError(t, err)

	return handler
}

func setRoutes(localizationHandler *localizator.Handler, endPointHandler *Handler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(localizationHandler.PassLocalizator)
	r.POST("/:lang/register", endPointHandler.CreateUser)

	return r
}

func setEnv() {
	os.Setenv("ENVIRONMENT", "TEST")
	os.Setenv("LOCALIZATOR_PATH", "/middlewares/localizator")
}

func TestCreateUserHandlerSuccess(t *testing.T) {
	setEnv()

	// DB INITIALIZATION
	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	var res response

	// SUCCESS USER 1
	payload := `{
		"nama":     "vwxyz",
		"noHp":     "+6283785332789",
		"email":    "admin@xyz.com",
		"password": "123456781234567812",
		"noRek":    "1234"
	}`
	req, err := http.NewRequest("POST", "/en/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)

	// SUCCESS USER 2
	payload = `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@abc.com",
		"password": "123456781234567812",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("POST", "/en/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
}

func TestCreateUserHandlerErrorBind(t *testing.T) {
	setEnv()

	// DB INITIALIZATION
	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	var res response

	// SUCCESS USER 1
	payload := `{
		"nama":     "vwxyz",
		"noHp":     "+6283785332789",
		"email":    "admin@xyz.com",
		"password": "123456781234567812",
		"noRek":    "1234"
	}`
	req, err := http.NewRequest("POST", "/en/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)

	// ERROR BIND USER 2
	payload = `{
		"nama":     "",
		"noHp":     "+6282876443890",
		"email":    "admin@abc.com",
		"password": "123456781234567812",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("POST", "/en/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Key: 'DataRequest.Nama' Error:Field validation for 'Nama' failed on the 'required' tag", res.Message)
}

func TestCreateUserHandlerErrorRequest(t *testing.T) {
	setEnv()

	// DB INITIALIZATION
	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	var res response

	// SUCCESS USER 1
	payload := `{
		"nama":     "vwxyz",
		"noHp":     "+6283785332789",
		"email":    "admin@xyz.com",
		"password": "123456781234567812",
		"noRek":    "1234"
	}`
	req, err := http.NewRequest("POST", "/en/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)

	// ERROR REQUEST USER 2
	payload = `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@xyz.com",
		"password": "123456781234567812",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("POST", "/en/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "badRequest", res.Message)
}
