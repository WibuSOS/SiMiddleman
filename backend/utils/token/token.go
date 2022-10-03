package token

import (
	"os"
	"time"

	"github.com/WibuSOS/sinarmas/models"
	//"github.com/WibuSOS/sinarmas/utils/errors"

	"github.com/dgrijalva/jwt-go"
)

const ExpTimeMinute = 60

func GenerateToken(user models.Users) (string, int64, error) {
	expTime := time.Now().Add(time.Minute * ExpTimeMinute).Unix()

	actClaims := jwt.MapClaims{}
	actClaims["user_id"] = user.ID
	actClaims["user_email"] = user.Email
	actClaims["exp"] = expTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, actClaims)
	resultToken, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", 0, err
	}

	return resultToken, expTime, nil
}
