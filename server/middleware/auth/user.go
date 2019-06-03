package auth

import (
	"server/middleware/jwt"
	"server/middleware/language"
	"server/model"
	"server/pkg/codes"
	"server/pkg/gdb"
	"errors"

	"github.com/gin-gonic/gin"
)

func searchUser(c *gin.Context) (user model.User, err error) {
	tokenString, isExist := c.Get("token")
	if !isExist {
		return user, errors.New(language.GetMsg(codes.ERROR_USER_NOT_FOUND))
	}
	claims, tokenErr := jwt.ParseToken(tokenString.(string))

	if tokenErr != nil {

	}

	userID := claims.UserId
	if err := gdb.Instance().Where("id = ?", userID).Find(&user).Error; err == nil {
		return user, nil
	}

	return user, errors.New(language.GetMsg(codes.ERROR_USER_NOT_FOUND))
}
