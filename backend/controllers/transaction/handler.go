package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) UpdateStatusDelivery(c *gin.Context) {
	id := c.Param("room_id")
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")
	var req RequestUpdateStatus

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": localizator.(*language.Config).Lookup(langReq, "badRequest")})
		return
	}

	err := h.Service.UpdateStatusDelivery(id, req)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": localizator.(*language.Config).Lookup(langReq, "successupdatestatus"),
	})
}

func (h *Handler) GetPaymentDetails(c *gin.Context) {
	idRoom := c.Param("room_id")
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")

	id, errConvert := strconv.Atoi(idRoom)
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": localizator.(*language.Config).Lookup(langReq, "badRequest")})
		return
	}

	res, err := h.Service.GetPaymentDetails(id)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": localizator.(*language.Config).Lookup(langReq, "successgetpaymentdetail"),
		"data":    res,
	})
}
