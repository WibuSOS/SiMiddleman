package authorization

import (
	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
)

type HandlerAuthorization struct {
	AuthorizationService ServiceAuthorization
}

func NewHandlerAuthorization(service ServiceAuthorization) *HandlerAuthorization {
	return &HandlerAuthorization{service}
}

func (h *HandlerAuthorization) RoleAuthorize(c *gin.Context) {
	localizator := c.MustGet("localizator")

	role := c.MustGet("role")
	roleStr := role.(string)

	err := h.AuthorizationService.RoleAuthorize(roleStr)
	if err != nil {
		msg := localizator.(*language.Config).Lookup(c.Param("lang"), err.Message)
		c.JSON(err.Status, gin.H{
			"message": msg,
		})
		c.Abort()
		return
	}

	c.Next()
}

func (h *HandlerAuthorization) RoomAuthorize(c *gin.Context) {
	localizator := c.MustGet("localizator")

	user_id := c.MustGet("id")
	userID := user_id.(float64)
	room_id := c.Param("room_id")

	err := h.AuthorizationService.RoomAuthorize(userID, room_id)
	if err != nil {
		msg := localizator.(*language.Config).Lookup(c.Param("lang"), err.Message)
		c.JSON(err.Status, gin.H{
			"message": msg,
		})
		c.Abort()
		return
	}

	c.Next()
}
