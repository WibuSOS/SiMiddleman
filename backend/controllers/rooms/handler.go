package rooms

import (
	"net/http"

	"github.com/WibuSOS/sinarmas/backend/utils/errors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req DataRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		error := errors.NewBadRequestError(err.Error())
		errors.LogError(error)
		c.JSON(error.Status, gin.H{"message": error.Message})
		return
	}

	newRoom, err := h.Service.CreateRoom(&req)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": err.Message,
			"data":    newRoom,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    newRoom,
	})
}

func (h *Handler) GetAllRooms(c *gin.Context) {
	userId := c.Param("id")
	newRooms, err := h.Service.GetAllRooms(userId)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": err.Message,
			"data":    newRooms,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    newRooms,
	})
}

func (h *Handler) JoinRoom(c *gin.Context) {
	roomId := c.Param("room_id")
	userId := c.Param("user_id")
	room, err := h.Service.JoinRoom(roomId, userId)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": err.Message,
		})
		return
	}

	statusArr := []string{"mulai transaksi", "barang dibayar", "barang dikirim", "konfirmasi barang sampai"}

	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"data":     room,
		"statuses": statusArr,
	})
}

func (h *Handler) JoinRoomPembeli(c *gin.Context) {
	roomId := c.Param("room_id")
	userId := c.Param("user_id")
	err := h.Service.JoinRoomPembeli(roomId, userId)
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
