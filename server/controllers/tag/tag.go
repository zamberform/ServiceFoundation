package tag

import (
	"log"
	"net/http"
	"server/controllers/error"
	"server/middleware/jwt"
	"server/models/database"
	"server/models/request"
	"server/models/response"
	"server/pkg/codes"
	"server/pkg/gdb"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	var commonRequest request.CommonReq
	if err := c.ShouldBindJSON(&commonRequest); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}
	userId := commonRequest.User.UserId
	uuid := commonRequest.User.UUID
	token := commonRequest.User.Token
	var user database.User
	if err := gdb.Instance().Where("id = ?", userId).Find(&user).Error; err == nil {
		userId = user.ID
	} else {
		var newUser database.User
		newUser.UUID = uuid
		newUser.CreatedAt = time.Now()
		newUser.UpdatedAt = time.Now()
		newUser.PlatformId = commonRequest.Platform

		if err := gdb.Instance().Create(&newUser).Error; err != nil {
			log.Fatalf("get.db.UserLoging err: %v", err)
			error.SendErrJSON("error", c)
			return
		}

		userId = user.ID
	}

	//check token is enable
	if _, err := jwt.ParseToken(token); err != nil {
		token, err = jwt.GenerateToken(string(userId), uuid)
		if err != nil {
			log.Fatalf("get.db.UserLoging err: %v", err)
			error.SendErrJSON("error", c)
			return
		}
	}

	var response response.CommonRes
	response.Msg = "success"
	response.Code = codes.SUCCESS

	response.User.Token = token
	c.JSON(http.StatusOK, response)
}

// For Admin Api
func AddTag(c *gin.Context) {

}

func UpdateTag(c *gin.Context) {

}

func DeleteTag(c *gin.Context) {

}
