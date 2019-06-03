package jwt

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}