package auth

import (
	"net/http"

	//"github.com/WibuSOS/sinarmas/utils/errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Login(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, token, err := h.Service.Login(req)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
		"token":   token,
	})
}
