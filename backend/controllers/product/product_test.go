package product

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/models"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func newTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	assert.NoError(t, err)
	assert.NotNil(t, db)

	err = db.AutoMigrate(&models.Users{}, &models.Products{}, &models.Rooms{})
	assert.NoError(t, err)

	db.Create(&models.Users{
		Nama:     "test2",
		NoHp:     "081234523415",
		NoRek:    "12341415",
		Email:    "test12@gmail.com",
		Password: "123456789",
	})

	db.Create(&models.Rooms{
		PenjualID: 1,
	})

	return db
}

func TestUpdateProduct(t *testing.T) {

	db := newTestDB(t)
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	db.Create(&models.Products{
		RoomsID:   1,
		Nama:      "produk1",
		Harga:     15000,
		Kuantitas: 2,
		Deskripsi: "ini adalah produk 1",
	})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.PUT("/updateproduct/:id", handler.UpdateProduct)
	payload := `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err := http.NewRequest("PUT", "/updateproduct/1", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	type response struct {
		Message string       `json:"message"`
		Data    DataResponse `json:"data"`
	}

	var res response
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "berhasil mengupdate data", res.Message)

	//error status bad request
	payload = `{"nama": "makanan dingin", "harga": "15000", "kuantitas": "2", "deskripsi":"ini adalah makanan dingin"}`
	req, err = http.NewRequest("PUT", "/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	//error Validate request name
	payload = `{"nama": "", "harga": 15000, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err = http.NewRequest("PUT", "/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	//error Validate request Harga
	payload = `{"nama": "makanan dingin", "harga": 0, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err = http.NewRequest("PUT", "/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	//error Validate request kuantitas
	payload = `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 0, "deskripsi":"ini adalah makanan dingin"}`
	req, err = http.NewRequest("PUT", "/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	//error Validate request deskripsi
	payload = `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 2, "deskripsi":""}`
	req, err = http.NewRequest("PUT", "/updateproduct/1", strings.NewReader(payload))
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)

	//error id not found
	payload = `{"nama": "makanan dingin", "harga": 15000, "kuantitas": 2, "deskripsi":"ini adalah makanan dingin"}`
	req, err = http.NewRequest("PUT", "/updateproduct/5", strings.NewReader(payload))
	assert.Nil(t, err)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)

	//error invalid id type
	// db.Create(&models.Products{
	// 	RoomsID:   1,
	// 	Nama:      "produk2",
	// 	Harga:     15000,
	// 	Kuantitas: 2,
	// 	Deskripsi: "ini adalah produk 2",
	// })

	// request := DataRequest{
	// 	Nama:      "produk2",
	// 	Harga:     15000,
	// 	Kuantitas: 2,
	// 	Deskripsi: "ini adalah produk 2",
	// }
	// _, error := repo.UpdateProduct("%", request)
	// assert.Nil(t, error)

	// w = httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// assert.Equal(t, 500, w.Code)

	// assert.Equal(t, "ini adalah makanan dingin", res.Data[0].Deskripsi)
	// assert.Equal(t, "makanan dingin", res.Data[0].Nama)

	// SUCCESS
	// payload := `{
	// "nama": "admin",
	// "email": "admin@admin.com",
	// "password": "123456781234567812",
	// "noHp": "+6281993220999",
	// "noRek": "1234"
	// 	}`
	// req, err := http.NewRequest("POST", "/register", strings.NewReader(payload))
	// assert.NoError(t, err)
	// assert.NotNil(t, req)

	// w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	// assert.Equal(t, "success", res.Message)
}
