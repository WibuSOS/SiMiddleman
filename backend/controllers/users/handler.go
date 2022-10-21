package users

import (
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

func (h *Handler) CreateUser(c *gin.Context) {
	var req DataRequest
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")
	if err := c.ShouldBindJSON(&req); err != nil {
		error := errors.NewBadRequestError(err.Error())
		errors.LogError(error)
		c.JSON(error.Status, gin.H{"message": error.Message})
		return
	}

	err := h.Service.CreateUser(&req)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": localizator.(*language.Config).Lookup(langReq, "success"),
	})
}

func (h *Handler) GetUserDetails(c *gin.Context) {
	idUser := c.Param("user_id")

	res, err := h.Service.GetUserDetails(idUser)
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

func (h *Handler) UpdateUser(c *gin.Context) {
	idUser := c.Param("user_id")
	var req DataRequestUpdateProfile
	if err := c.ShouldBindJSON(&req); err != nil {
		error := errors.NewBadRequestError(err.Error())
		errors.LogError(error)
		c.JSON(error.Status, gin.H{"message": error.Message})
		return
	}

	err := h.Service.UpdateUser(idUser, req)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
