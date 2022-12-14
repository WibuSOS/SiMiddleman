package product

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
	r.PUT("/:lang/updateproduct/:product_id", endPointHandler.UpdateProduct)

	return r
}

func setEnv() {
	os.Setenv("ENVIRONMENT", "TEST")
	os.Setenv("LOCALIZATOR_PATH", "")
}

func TestUpdateProduct(t *testing.T) {
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

	payload := `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err := http.NewRequest("PUT", "/en/updateproduct/1", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Success update data product", res.Message)
}

func TestErrorStatusBadRequest(t *testing.T) {
	setEnv()

	// DB INITIALIZATION
	db := newTestDB(t)

	var res response

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)
	//error status bad request

	payload := `{"nama": "makanan dingin", "harga": "15000", "kuantitas": "2", "deskripsi":"ini adalah makanan dingin"}`
	req, err := http.NewRequest("PUT", "/en/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Bad Request", res.Message)
}

func TestErrorNameEmpty(t *testing.T) {
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

	//error Validate request name
	payload := `{"nama": "", "harga": 15000, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err := http.NewRequest("PUT", "/en/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Name Cannot Be Empty", res.Message)
}

func TestErrorPriceEmpty(t *testing.T) {
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

	//error Validate request Harga
	payload := `{"nama": "makanan dingin", "harga": 0, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err := http.NewRequest("PUT", "/en/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Price Cannot Be Empty", res.Message)
}

func TestErrorQuantityEmpty(t *testing.T) {
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

	//error Validate request kuantitas
	payload := `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 0, "deskripsi":"ini adalah makanan dingin"}`
	req, err := http.NewRequest("PUT", "/en/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Quantity Cannot Be Empty", res.Message)

}

func TestErrorDescriptionEmpty(t *testing.T) {
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

	//error Validate request deskripsi
	payload := `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 2, "deskripsi":""}`
	req, err := http.NewRequest("PUT", "/en/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Description Cannot Be Empty", res.Message)
}

func TestErrorIdNotFound(t *testing.T) {
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

	//error id not found
	payload := `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err := http.NewRequest("PUT", "/en/updateproduct/5", strings.NewReader(payload))
	assert.Nil(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Record Not Found", res.Message)
}
