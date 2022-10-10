package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentDetailsHandlerSuccess(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	idRoom := newTestCreateRoomWithProduct(t)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getHarga/:idroom", handler.GetPaymentDetails)

	url := fmt.Sprintf("/getHarga/%d", idRoom)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type response struct {
		Message string              `json:"message"`
		Data    ResponsePaymentInfo `json:"data"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Greater(t, res.Data.Total, 0)
}

func TestGetPaymentDetailsHandlerEmptyProduct(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	idRoom := newTestCreateRoomWithoutProduct(t)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getHarga/:idroom", handler.GetPaymentDetails)

	url := fmt.Sprintf("/getHarga/%d", idRoom)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type response struct {
		Message string              `json:"message"`
		Data    ResponsePaymentInfo `json:"data"`
	}
	var res response

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, 0, res.Data.Total)
}

func TestGetPaymentDetailsHandlerRoomNotFound(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/getHarga/:idroom", handler.GetPaymentDetails)

	url := fmt.Sprintf("/getHarga/%d", 3)
	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
