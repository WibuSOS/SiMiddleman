package authorization

import (
	"strconv"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type roomAuth struct {
	db *gorm.DB
}

func NewRoomAuth(db *gorm.DB) *roomAuth {
	return &roomAuth{db}
}

func (r *roomAuth) RoomAuthorize(c *gin.Context) {
	var room models.Rooms

	user_id := c.MustGet("id")
	userID := user_id.(float64)
	room_id := c.Param("room_id")

	roomID, err := strconv.ParseUint(room_id, 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid Room ID")
		c.JSON(restErr.Status, gin.H{
			"message": restErr.Message,
		})
		c.Abort()
		return
	}

	res := r.db.Where("id = ?", roomID).First(&room)
	if res.Error != nil {
		restErr := errors.NewBadRequestError("Room not found")
		c.JSON(restErr.Status, gin.H{
			"message": restErr.Message,
		})
		c.Abort()
		return
	}

	isPenjual := uint(userID) == room.PenjualID
	var isPembeli bool
	if room.PembeliID == nil {
		isPembeli = false
	} else {
		isPembeli = uint(userID) == *room.PembeliID
	}

	if !(isPenjual || isPembeli) {
		restErr := errors.NewUnauthorized("Unauthorized")
		c.JSON(restErr.Status, gin.H{
			"message": restErr.Message,
		})
		c.Abort()
		return
	}

	c.Next()
}
