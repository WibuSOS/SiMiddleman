package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandlerSuccess(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}
	var res response

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", handler.CreateUser)

	// SUCCESS USER 1
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

	// SUCCESS USER 2
	payload = `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@abc.com",
		"password": "123456781234567812",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
}

func TestCreateUserHandlerErrorBind(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}
	var res response

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", handler.CreateUser)

	// SUCCESS USER 1
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

	// ERROR BIND USER 2
	payload = `{
		"nama":     "",
		"noHp":     "+6282876443890",
		"email":    "admin@abc.com",
		"password": "123456781234567812",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Key: 'DataRequest.Nama' Error:Field validation for 'Nama' failed on the 'required' tag", res.Message)
}

func TestCreateUserHandlerErrorRequest(t *testing.T) {
	type response struct {
		Message string `json:"message"`
	}
	var res response

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", handler.CreateUser)

	// SUCCESS USER 1
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

	// ERROR REQUEST USER 2
	payload = `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@xyz.com",
		"password": "123456781234567812",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("POST", "/register", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "constraint failed: UNIQUE constraint failed: users.email (2067)", res.Message)
}
