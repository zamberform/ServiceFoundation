package user

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

func Login(c *gin.Context) {
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

func SignIn(c *gin.Context) {
	var commonReq request.CommonReq
	commonInfo, _ := c.Get("common")

	commonReq = commonInfo.(request.CommonReq)

	userId := commonReq.User.UserId
	token := commonReq.User.Token
	var user database.User
	if err := gdb.Instance().Where("id = ?", userId).Find(&user).Error; err == nil {
		userId = user.ID
	}

	user.Introduce = "sign_in"
	user.Email = "brightzamber@gmail.com"
	if err := gdb.Instance().Model(&database.User{}).Where("id = ?", userId).Updates(user).Error; err != nil {
		log.Fatalf("get.db.Signin err: %v", err)
		return
	}

	var response response.CommonRes
	response.Msg = "success"
	response.Code = codes.SUCCESS

	response.User.Token = token
	c.JSON(http.StatusOK, response)
}

func Withdrawal(c *gin.Context) {
	var commonReq request.CommonReq
	commonInfo, _ := c.Get("common")

	commonReq = commonInfo.(request.CommonReq)

	userId := commonReq.User.UserId
	token := commonReq.User.Token
	var user database.User
	if err := gdb.Instance().Where("id = ?", userId).Find(&user).Error; err == nil {
		userId = user.ID
	}

	user.Introduce = "quit"
	if err := gdb.Instance().Model(&database.User{}).Where("id = ?", userId).Updates(user).Error; err != nil {
		log.Fatalf("get.db.Signin err: %v", err)
		return
	}

	var response response.CommonRes
	response.Msg = "success"
	response.Code = codes.SUCCESS

	response.User.Token = token
	c.JSON(http.StatusOK, response)
}

// For Admin Api
func AddUser(c *gin.Context) {

}

func GetUserList(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
