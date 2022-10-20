package rooms

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

func (h *Handler) CreateRoom(c *gin.Context) {
	var req DataRequest
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")
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
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
			"data":    newRoom,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": localizator.(*language.Config).Lookup(langReq, "successcreateroom"),
		"data":    newRoom,
	})
}

func (h *Handler) GetAllRooms(c *gin.Context) {
	userId := c.Param("id")
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")
	newRooms, err := h.Service.GetAllRooms(userId)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
			"data":    newRooms,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": localizator.(*language.Config).Lookup(langReq, "successgetallroom"),
		"data":    newRooms,
	})
}

func (h *Handler) JoinRoom(c *gin.Context) {
	roomId := c.Param("room_id")
	userId := c.Param("user_id")
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")

	room, err := h.Service.JoinRoom(roomId, userId)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
		})
		return
	}

	statusArr := []string{"mulai transaksi", "barang dibayar", "barang dikirim", "konfirmasi barang sampai"}

	c.JSON(http.StatusOK, gin.H{
		"message":  localizator.(*language.Config).Lookup(langReq, "successjoinroom"),
		"data":     room,
		"statuses": statusArr,
	})
}

func (h *Handler) JoinRoomPembeli(c *gin.Context) {
	roomId := c.Param("room_id")
	userId := c.Param("user_id")
	langReq := c.Param("lang")
	localizator := c.MustGet("localizator")
	err := h.Service.JoinRoomPembeli(roomId, userId)
	if err != nil {
		errors.LogError(err)
		c.JSON(err.Status, gin.H{
			"message": localizator.(*language.Config).Lookup(langReq, err.Message),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": localizator.(*language.Config).Lookup(langReq, "successjoinroombuyer"),
	})
}
