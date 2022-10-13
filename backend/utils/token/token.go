package token

import (
	"log"
	"os"
	"time"

	"github.com/WibuSOS/sinarmas/backend/models"
	"github.com/WibuSOS/sinarmas/backend/utils/errors"

	"github.com/dgrijalva/jwt-go"
)

const ExpTimeMinute = 60

func GenerateToken(user models.Users) (string, *errors.RestError) {
	expTime := time.Now().Add(time.Minute * ExpTimeMinute).Unix()

	actClaims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, actClaims)
	resultToken, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("Login: Error while creating token")
		return "", errors.NewInternalServerError("Error while creating token")
	}

	return resultToken, nil
}

func ValidationToken(dataTime int64) bool {
	if dataTime > time.Now().Unix() {
		log.Println("Token Validation: Token Not Expired")
		return false
	}

	log.Println("Token Validation: Expired")
	return true
}
