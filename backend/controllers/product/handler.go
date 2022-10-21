package product

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

func (h *Handler) UpdateProduct(c *gin.Context) {
	var req DataRequest
	id := c.Param("product_id")
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")

	if err := c.ShouldBindJSON(&req); err != nil {
		errRest := errors.NewBadRequestError(err.Error())
		log.Println("Status Bad Request : ", err)
		c.JSON(errRest.Status, gin.H{"message": localizator.(*language.Config).Lookup(langReq, "badRequest")})
		return
	}

	res, err := h.Service.UpdateProduct(id, req)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": localizator.(*language.Config).Lookup(langReq, "successupdate"),
		"Data":    res,
	})
}
