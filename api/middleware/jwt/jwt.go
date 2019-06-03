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

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		utils.EncodeMD5(username),
		utils.EncodeMD5(password),
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
