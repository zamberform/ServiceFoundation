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
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	User     string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var currentUser User
	if err := c.ShouldBindJSON(&currentUser); err != nil {
		error.SendErrJSON("error", c)
		return
	}
	realUser := database.User{}
	if err := gdb.Instance().First(&realUser, "name = ?", currentUser.User).Error; err != nil {
		log.Printf("find User err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	//パスワード比較
	if err := bcrypt.CompareHashAndPassword([]byte(realUser.Pass), []byte(currentUser.Password)); err != nil {
		log.Printf("get.db.Signin err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if token, err := jwt.GenerateToken(realUser.Name, realUser.Pass); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"user": gin.H{
				"name":  realUser.Name,
				"token": token,
				"limit": 1,
			},
		})
	} else {
		log.Printf("get.db.UserLoging err: %v", err)
		error.SendErrJSON("error", c)
	}
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

func CheckAuth(c *gin.Context) {
	if values, _ := c.Request.Header["Auth-Token"]; len(values) > 0 {
		userToken := values[0]
		loginInfo, err := jwt.ParseToken(userToken)
		if err != nil {
			log.Printf("get.db.UserLoging err: %v", err)
			error.SendErrJSON("error", c)
		} else {
			loginUser := database.User{}
			if err := gdb.Instance().First(&loginUser, "name = ?", loginInfo.UserId).Error; err != nil {
				log.Printf("get.db.User.Login err: %v", err)
				error.SendErrJSON("error", c)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"msg":    "success",
			})
		}
	}
}

func Register(c *gin.Context) {
	var addUser database.User
	if err := c.ShouldBindJSON(&addUser); err != nil {
		error.SendErrJSON("error", c)
		return
	}
	addUser.PlatformId = 1
	hash, err := bcrypt.GenerateFromPassword([]byte(addUser.Pass), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	var user database.User
	var isExist = false
	if err := gdb.Instance().Where("name = ?", addUser.Name).First(&user).Error; err == nil {
		isExist = true
	}

	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"status": 300,
			"msg":    "存在しているユーザー",
		})
		return
	}
	addUser.Pass = string(hash)
	addUser.CreatedAt = time.Now()
	addUser.UpdatedAt = time.Now()
	if err := gdb.Instance().Create(&addUser).Error; err != nil {
		log.Printf("find user err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})
}

func Logout(c *gin.Context) {
	buf := make([]byte, 2048)
	n, _ := c.Request.Body.Read(buf)
	userToken := string(buf[0:n])
	blackList := database.TokenBlackList{}
	blackList.Token = userToken
	if err := gdb.Instance().Create(&blackList).Error; err != nil {
		log.Printf("logout failed err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})

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
	addUser.PlatformId = 1
	hash, err := bcrypt.GenerateFromPassword([]byte(addUser.Pass), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	var user database.User
	var isExist = false
	if err := gdb.Instance().Where("name = ?", addUser.Name).First(&user).Error; err == nil {
		isExist = true
	}

	if isExist {
		c.JSON(http.StatusOK, gin.H{
			"status": 300,
			"msg":    "存在しているユーザー",
		})
		return
	}
	addUser.Pass = string(hash)
	addUser.CreatedAt = time.Now()
	addUser.UpdatedAt = time.Now()
	if err := gdb.Instance().Create(&addUser).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
	})
}

func GetUserList(c *gin.Context) {
	users := []database.User{}
	if err := gdb.Instance().Find(&users).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
		"users":  users,
	})
}

func UpdateUserDesc(c *gin.Context) {
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
	updateAfterUser.Introduce = updateInfo.Introduce
	if err := gdb.Instance().Model(&updateBeforeUser).Update(&updateAfterUser).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	if err := gdb.Instance().Save(&updateAfterUser).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
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
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	delUser.ID = uint(userId)

	if err := gdb.Instance().Find(&delUser).Error; err != nil {
		log.Printf("not find user err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Delete(&delUser).Error; err != nil {
		log.Printf("delete user fail err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
	})
}
