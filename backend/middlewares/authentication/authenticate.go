package authentication

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/WibuSOS/sinarmas/backend/utils/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	language "github.com/moemoe89/go-localization"
)

func Authentication(c *gin.Context) {
	localizator := c.MustGet("localizator")
	tokenString := c.Request.Header.Get("Authorization")
	str := tokenString
	str = strings.ReplaceAll(str, "Bearer ", "")
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("can't verify token")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println(err)
		msg := localizator.(*language.Config).Lookup(c.Param("lang"), "noToken")
		restErr := errors.NewUnauthorized(msg)
		c.JSON(restErr.Status, gin.H{
			"message": restErr.Message,
		})
		c.Abort()
		return
	}
	log.Println("Authentication: Token Verified!")

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		msg := localizator.(*language.Config).Lookup(c.Param("lang"), "invalidToken")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		c.Abort()
		return
	}

	id := claims["ID"]
	role := claims["Role"]

	c.Set("id", id)
	c.Set("role", role)
	c.Next()
}
