package jwt

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserId string `json:"userId"`
	UUID   string `json:"uuid"`
	jwt.StandardClaims
}
