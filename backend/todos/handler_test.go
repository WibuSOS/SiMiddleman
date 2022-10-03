package todos

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/send", handler.CreateTodo)
	payload := `{"task": "task 1"}`
	req, err := http.NewRequest("POST", "/send", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string       `json:"message"`
		Data    models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, "task 1", res.Data.Task)
}

func TestErrorBindCreateTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/send", handler.CreateTodo)
	payload := `{"task": 1}`
	req, err := http.NewRequest("POST", "/send", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	type response struct {
		Error string `json:"error"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "json: cannot unmarshal number into Go struct field DataRequest.task of type string", res.Error)
}

func TestErrorInsertCreateTodoHandler(t *testing.T) {
	db := newEmptyTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/send", handler.CreateTodo)
	payload := `{"task": "task 1"}`
	req, err := http.NewRequest("POST", "/send", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	type response struct {
		Message string       `json:"message"`
		Data    models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "SQL logic error: no such table: todos (1)", res.Message)
	assert.Equal(t, models.Todos{}, res.Data)
}

func TestGetTodoHandler(t *testing.T) {
	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	db.Create(&models.Todos{
		Task: "task 1",
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", handler.GetTodos)
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	type response struct {
		Message string         `json:"message"`
		Data    []models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "success", res.Message)
	assert.Equal(t, "task 1", res.Data[0].Task)
}

func TestErrorGetTodoHandler(t *testing.T) {
	db := newEmptyTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", handler.GetTodos)
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	type response struct {
		Message string         `json:"message"`
		Data    []models.Todos `json:"data"`
	}
	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Len(t, res.Data, 0)
}
