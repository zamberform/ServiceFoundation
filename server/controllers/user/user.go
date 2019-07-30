package user

import (
	"fmt"
	"log"
	"net/http"
	"server/controllers/error"
	"server/middleware/jwt"
	"server/models/database"
	"server/models/request"
	"server/models/response"
	"server/pkg/codes"
	"server/pkg/gdb"
	"strconv"
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
	addUser := database.User{}
	if err := c.ShouldBindJSON(&addUser); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}

	if err := gdb.Instance().Create(&addUser).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, "success")
}

func GetUserList(c *gin.Context) {
	users := []database.User{}
	if err := gdb.Instance().Find(&users).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, "success")
}

func UpdateUser(c *gin.Context) {
	updateInfo := database.User{}
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}

	userIdStr := c.Param("id")
	updateBeforeUser := database.User{}
	// 削除したいレコードのIDを指定
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	updateBeforeUser.ID = uint(userId)

	if err := gdb.Instance().Find(&updateBeforeUser).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	updateAfterUser := updateBeforeUser
	updateAfterUser.Name = updateInfo.Name
	updateAfterUser.Email = updateInfo.Email
	updateAfterUser.Introduce = updateInfo.Introduce
	updateAfterUser.AvatarURL = updateInfo.AvatarURL
	if err := gdb.Instance().Model(&updateBeforeUser).Update(&updateAfterUser).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	if err := gdb.Instance().Save(&updateAfterUser).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	c.JSON(http.StatusOK, "success")
}

func DeleteUser(c *gin.Context) {
	userIdStr := c.Param("id")
	delUser := database.User{}
	// 削除したいレコードのIDを指定
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	delUser.ID = uint(userId)

	if err := gdb.Instance().Find(&delUser).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Delete(&delUser).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	c.JSON(http.StatusOK, "success")
}
