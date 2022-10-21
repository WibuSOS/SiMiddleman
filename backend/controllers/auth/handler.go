package auth

import (
	"log"
	"net/http"

	"github.com/WibuSOS/sinarmas/backend/utils/errors"

	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Login(c *gin.Context) {
	localizator := c.MustGet("localizator")

	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println(err.Error())
		msg := localizator.(*language.Config).Lookup(c.Param("lang"), "invalidJSON")
		c.JSON(http.StatusBadRequest, gin.H{"message": msg})
		return
	}

	res, token, err := h.Service.Login(req)
	if err != nil {
		errors.LogError(err)
		msg := localizator.(*language.Config).Lookup(c.Param("lang"), err.Message)
		c.JSON(err.Status, gin.H{
			"message": msg,
		})
		return
	}

	msg := localizator.(*language.Config).Lookup(c.Param("lang"), "success")
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"data":    res,
		"token":   token,
	})
}
