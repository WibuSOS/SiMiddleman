package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/middlewares/localizator"
	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type responseGetUserDetails struct {
	Message string       `json:"message"`
	Data    models.Users `json:"data"`
}

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
	r.GET("/:lang/user/:user_id", endPointHandler.GetUserDetails)
	r.PUT("/:lang/user/:user_id", endPointHandler.UpdateUser)

	return r
}

func setEnv() {
	os.Setenv("ENVIRONMENT", "TEST")
}

func TestMain(m *testing.M) {
	setEnv()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestCreateUserHandlerSuccess(t *testing.T) {
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
	assert.Equal(t, "Bad Request", res.Message)
}

func newTestGetUserDetailsHandler(t *testing.T, isError bool) (*httptest.ResponseRecorder, responseGetUserDetails) {
	var resGetDetails responseGetUserDetails
	var req *http.Request
	var err error

	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	if isError {
		req, err = http.NewRequest("GET", "/en/user/10", nil)
	} else {
		req, err = http.NewRequest("GET", "/en/user/1", nil)
	}

	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w, resGetDetails
}

func TestGetUserDetailsHandlerSuccess(t *testing.T) {
	w, userDetails := newTestGetUserDetailsHandler(t, false)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &userDetails))
	assert.Equal(t, "success", userDetails.Message)
}

func TestGetUserDetailsHandlerError(t *testing.T) {
	w, userDetails := newTestGetUserDetailsHandler(t, true)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &userDetails))
	assert.Equal(t, "Error while fetching data", userDetails.Message)
}

func TestUpdateUserHandlerSuccessRequest(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}
	var res response

	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	// UPDATE USER
	payload := `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@xyz.com",
		"noRek":    "6789"
	}`
	req, err := http.NewRequest("PUT", "/en/user/1", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
}

func TestUpdateUserHandlerErrorRequest(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}
	var res response

	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	// UPDATE USER
	req, err := http.NewRequest("PUT", "/en/user/1", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "invalid request", res.Message)
}

func TestUpdateUserHandlerErrorIDRequest(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}
	var res response

	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	// UPDATE USER
	payload := `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@xyz.com",
		"noRek":    "6789"
	}`
	req, err := http.NewRequest("PUT", "/en/user/1000", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "bad request", res.Message)
}
