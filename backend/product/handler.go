package product

import (
	"log"
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

func (h *Handler) GetSpesifikProduct(c *gin.Context) {
	var req DataRequest
	idroom := c.Param("idroom")

	res, err := h.Service.GetSpesifikProduct(idroom, req)
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

func (h *Handler) CreateProduct(c *gin.Context) {
	var req DataRequest
	idroom := c.Param("idroom")
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Status Bad Request : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateProduct(idroom, req)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success Create Product",
		"Data":    res,
	})
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	var req DataRequest
	id := c.Param("id")
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Status Bad Request : ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.UpdateProduct(id, req)
	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil mengupdate data",
		"Data":    res,
	})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := h.Service.DeleteProduct(id)

	if err != nil {
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "berhasil menghapus Data!",
	})
}
