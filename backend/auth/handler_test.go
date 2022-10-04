package auth

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	//"github.com/WibuSOS/sinarmas/models"
// 	//"github.com/WibuSOS/sinarmas/utils/errors"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestLoginHandler(t *testing.T) {
// 	db := newTestDB(t)
// 	repo := NewRepository(db)
// 	service := NewService(repo)
// 	handler := NewHandler(service)

// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default()
// 	r.POST("/login", handler.Login)
// 	payload := `{"email": "fikri@gmail.com", "password": "fikri123"}`
// 	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
// 	assert.NoError(t, err)
// 	assert.NotNil(t, req)

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var res DataResponse
// 	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
// }
