package transaction

import (
	"net/http"

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
