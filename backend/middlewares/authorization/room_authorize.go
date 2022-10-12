package authorization

import (
	"github.com/WibuSOS/sinarmas/utils/errors"

	"github.com/gin-gonic/gin"
)

func (r Roles) RoomAuthorize(c *gin.Context) {
	//userID := c.MustGet("ID")

	restErr := errors.NewUnauthorized("Unauthorized")
	c.JSON(restErr.Status, gin.H{
		"message": restErr.Message,
	})
	c.Abort()
}
