package user

import (
	"api/middleware/jwt"
	"api/model"
	"api/pkg/gdb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	uuid, _ := c.Get("uuid")
	platformId, _ := c.Get("platform")

	var user model.User
	var userId uint
	var userPlatformString string
	if err := gdb.Instance().Where("uuid = ? And platform = ?", uuid, platformId).Find(&user).Error; err == nil {
		userId = user.ID
		userPlatformString = user.Pass
	} else {
		var newUser model.User
		newUser.UUID = uuid.(string)
		newUser.Platform = platformId.(int)

		if err := gdb.Instance().Create(&newUser).Error; err != nil {

		}

		userId = user.ID
		userPlatformString = user.UUID
	}

	token, err := jwt.GenerateToken(userId, userPlatformString)
	if err != nil {

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"token": token,
		},
	})
}

func SignIn(c *gin.Context) {

}

func Withdrawal(c *gin.Context) {

}