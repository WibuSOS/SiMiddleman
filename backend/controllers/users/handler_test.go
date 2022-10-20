package users

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type responseGetUserDetails struct {
	Message string       `json:"message"`
	Data    models.Users `json:"data"`
}

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

func newTestGetUserDetailsHandler(t *testing.T, isError bool) (*httptest.ResponseRecorder, responseGetUserDetails) {
	type response struct {
		Message string `json:"message"`
	}
	var res response
	var resGetDetails responseGetUserDetails

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	dbError := newTestDBError(t)
	repoError := NewRepository(dbError)
	serviceError := NewService(repoError)
	handlerError := NewHandler(serviceError)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", handler.CreateUser)
	if isError {
		r.GET("/:lang/user/:user_id", handlerError.GetUserDetails)
	} else {
		r.GET("/:lang/user/:user_id", handler.GetUserDetails)
	}
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

	req, err = http.NewRequest("GET", "/en/user/1", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
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
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", handler.CreateUser)
	r.PUT("/user/:user_id", handler.UpdateUser)

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

	// UPDATE USER
	payload = `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@xyz.com",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("PUT", "/user/1", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
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
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", handler.CreateUser)
	r.PUT("/user/:user_id", handler.UpdateUser)

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

	req, err = http.NewRequest("PUT", "/user/1", nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
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
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/register", handler.CreateUser)
	r.PUT("/user/:user_id", handler.UpdateUser)

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

	// UPDATE USER
	payload = `{
		"nama":     "abcde",
		"noHp":     "+6282876443890",
		"email":    "admin@xyz.com",
		"noRek":    "6789"
	}`
	req, err = http.NewRequest("PUT", "/user/1000", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "bad request", res.Message)
}
