package jwt

import (
	"api/middleware/language"
	"api/pkg/codes"
	"api/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret []byte

func Init() {

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func GenerateToken(userId, uuid string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		utils.EncodeMD5(userId),
		utils.EncodeMD5(uuid),
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "service-api",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ApiJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = codes.SUCCESS
		token := c.Query("token")

		if token == "" {
			code = codes.INVALID_PARAMS
		} else {
			_, err := parseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = codes.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = codes.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != codes.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  language.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
