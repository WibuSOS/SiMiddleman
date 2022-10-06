package authorization

import (
	"github.com/WibuSOS/sinarmas/utils/errors"

	"github.com/gin-gonic/gin"
)

type Roles struct {
	AllowedRoles []string
}

func (r Roles) Authorize(c *gin.Context) {
	role := c.MustGet("role")
	for i := 0; i < len(r.AllowedRoles); i++ {
		if r.AllowedRoles[i] == role {
			c.Next()
			return
		}
	}

	restErr := errors.NewUnauthorized("Unauthorized")
	c.JSON(restErr.Status, gin.H{
		"message": restErr.Message,
	})
	c.Abort()
}
