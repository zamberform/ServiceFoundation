package user

import (
	"log"
	"net/http"
	"server/controllers/error"
	"server/middleware/jwt"
	"server/models/db"
	"server/models/request"
	"server/pkg/gdb"
	"server/pkg/setting"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var commonRequest request.CommonReq
	if err := c.ShouldBindJSON(&commonRequest); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}
	uuid := commonRequest.User.Uuid
	platformId := commonRequest.Platform
	var user db.User
	var userId uint
	var userPlatformString string
	if err := gdb.Instance().Where("uuid = ? And fk_platform_id = ?", uuid, platformId).Find(&user).Error; err == nil {
		userId = user.ID
		userPlatformString = user.Pass
	} else {
		var newUser db.User
		newUser.UUID = uuid

		if err := gdb.Instance().Create(&newUser).Error; err != nil {
			log.Fatalf("get.db.UserLoging err: %v", err)
			error.SendErrJSON("error", c)
			return
		}

		userId = user.ID
		userPlatformString = user.UUID
	}

	token, err := jwt.GenerateToken(string(userId), userPlatformString, setting.AppSetting.JwtSecret)
	if err != nil {
		log.Fatalf("get.db.UserLoging err: %v", err)
		error.SendErrJSON("error", c)
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
