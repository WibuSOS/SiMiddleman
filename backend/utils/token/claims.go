package token

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	ID   uint   `json:"ID"`
	Role string `json:"Role"`
}
