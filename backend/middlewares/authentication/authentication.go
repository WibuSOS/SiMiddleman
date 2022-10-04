package authentication

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/WibuSOS/sinarmas/utils/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	str := tokenString
	str = strings.ReplaceAll(str, "Bearer ", "")
	_, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("can't verify token")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println("Authentication: Unable to verify Token")
		restErr := errors.NewUnauthorized("Unable to verify Token")
		c.JSON(restErr.Status, restErr)
		c.Abort()
		return
	}

	log.Println("Authentication: Token Verified!")
}
