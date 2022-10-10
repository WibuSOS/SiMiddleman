package transaction

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) UpdateStatusDelivery(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.UpdateStatusDelivery(id)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success update status pengiriman barang",
	})
}

func (h *Handler) GetPaymentDetails(c *gin.Context) {
	idRoom := c.Param("idroom")
	id, errConvert := strconv.Atoi(idRoom)
	if errConvert != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errConvert.Error()})
		return
	}

	res, err := h.Service.GetPaymentDetails(id)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
}
