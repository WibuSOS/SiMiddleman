package transaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"
	"github.com/WibuSOS/sinarmas/backend/middlewares/localizator"
	"github.com/WibuSOS/sinarmas/backend/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type createRoomResponse struct {
	Message string       `json:"message"`
	Data    models.Rooms `json:"data"`
}

type getPaymentDetailsResponse struct {
	Message string              `json:"message"`
	Data    ResponsePaymentInfo `json:"data"`
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
	r.PUT("/:lang/rooms/details/updatestatus/:room_id", endPointHandler.UpdateStatusDelivery)
	r.GET("/:lang/rooms/details/getHarga/:room_id", endPointHandler.GetPaymentDetails)
	return r
}

func setEnv() {
	os.Setenv("ENVIRONMENT", "TEST")
	os.Setenv("LOCALIZATOR_PATH", "")
}

func TestUpdateStatusDelivery(t *testing.T) {
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

	payload := `{"status": "barang dibayar"}`
	req, err := http.NewRequest("PUT", "/en/rooms/details/updatestatus/1", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "Success Update Status", res.Message)

}

func TestUpdateStatusDeliveryInvalidJSON(t *testing.T) {
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

	payload := `{"status": 123}`
	req, err := http.NewRequest("PUT", "/en/rooms/details/updatestatus/1", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
}

func TestUpdateErrorStatusDelivery(t *testing.T) {

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

	payload := `{"status": "barang dibayar"}`
	req, err := http.NewRequest("PUT", "/en/rooms/details/updatestatus/asasdaweq", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
	assert.Equal(t, "WHERE conditions required", res.Message)
}

func newTestGetPaymentDetailsHandler(t *testing.T, withID bool, roomID *string) (*httptest.ResponseRecorder, getPaymentDetailsResponse) {
	var createRoomRes createRoomResponse
	var res getPaymentDetailsResponse

	setEnv()

	// DB INITIALIZATION
	db := newTestDB(t)

	// LOCALIZATION HANDLER
	localizationHandler := setLocalizationHandler(t)

	// END-POINT HANDLER
	endPointHandler := setEndPointHandler(t, db)

	// ROUTES INITIALIZATION
	r := setRoutes(localizationHandler, endPointHandler)

	roomRepo := rooms.NewRepository(db)
	roomService := rooms.NewService(roomRepo)
	roomHandler := rooms.NewHandler(roomService)

	r.POST("/:lang/rooms", roomHandler.CreateRoom)

	payload := `{
		"id": 1,
		"product" : {
			"nama": "Razer Mouse",
			"deskripsi": "Ini Razer Mouse",
			"harga": 150000,
			"kuantitas": 1
		}
	}`

	req, err := http.NewRequest("POST", "/en/rooms", strings.NewReader(payload))
	assert.NoError(t, err)
	assert.NotEmpty(t, req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &createRoomRes))
	assert.Equal(t, "Success Create Room, please refresh to view room", createRoomRes.Message)
	assert.NotEmpty(t, createRoomRes.Data.RoomCode)

	var url string
	if withID {
		url = fmt.Sprintf("/en/rooms/details/getHarga/%d", createRoomRes.Data.ID)
	} else {
		if roomID != nil {
			url = fmt.Sprintf("/en/rooms/details/getHarga/%s", *roomID)
		} else {
			url = fmt.Sprintf("/en/rooms/details/getHarga/%s", createRoomRes.Data.RoomCode)
		}
	}

	req, err = http.NewRequest("GET", url, nil)
	assert.NoError(t, err)
	assert.NotNil(t, req)

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w, res
}

func TestGetPaymentDetailsHandlerSuccess(t *testing.T) {
	w, paymentDetails := newTestGetPaymentDetailsHandler(t, true, nil)

	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &paymentDetails))
	assert.Equal(t, "Success get payment detail", paymentDetails.Message)
	assert.Greater(t, int(paymentDetails.Data.Total), 0)
}

func TestGetPaymentDetailsHandlerErrorQueryParam(t *testing.T) {
	w, _ := newTestGetPaymentDetailsHandler(t, false, nil)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetPaymentDetailsHandlerRoomNotFound(t *testing.T) {
	roomID := "7"
	w, _ := newTestGetPaymentDetailsHandler(t, false, &roomID)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
