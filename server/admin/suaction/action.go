package suaction

import (
	"log"
	"net/http"
	"server/controllers/error"
	"server/middleware/jwt"
	"server/models/database"
	"server/pkg/gdb"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminUser struct {
	User     string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var currentUser AdminUser
	if err := c.ShouldBindJSON(&currentUser); err != nil {
		error.SendErrJSON("error", c)
		return
	}
	adminUser := database.Admin{}
	if err := gdb.Instance().First(&adminUser, "name = ?", currentUser.User).Error; err != nil {
		log.Printf("get.db.Signin err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	//パスワード比較
	// if err := bcrypt.CompareHashAndPassword([]byte(adminUser.Pass), []byte(currentUser.Password)); err != nil {
	// 	log.Printf("get.db.Signin err: %v", err)
	// 	error.SendErrJSON("error", c)
	// 	return
	// }

	if token, err := jwt.GenerateToken(adminUser.Name, adminUser.Pass+string(time.Now().Unix())); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "success",
			"auth": gin.H{
				"token": token,
			},
		})
	} else {
		log.Fatalf("get.db.UserLoging err: %v", err)
		error.SendErrJSON("error", c)
	}
}

func Logout(c *gin.Context) {
	if values, _ := c.Request.Header["X-Token"]; len(values) > 0 {
		userToken := values[0]
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
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 300,
			"msg":    "failed",
		})
	}

}

func GetAdminInfo(c *gin.Context) {
	if values, _ := c.Request.Header["X-Token"]; len(values) > 0 {
		userToken := values[0]
		adminInfo, err := jwt.ParseToken(userToken)
		if err != nil {
			log.Fatalf("get.db.UserLoging err: %v", err)
			error.SendErrJSON("error", c)
		} else {
			adminUser := database.Admin{}
			if err := gdb.Instance().First(&adminUser, "name = ?", adminInfo.UserId).Error; err != nil {
				log.Printf("get.db.Signin err: %v", err)
				error.SendErrJSON("error", c)
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"msg":    "success",
				"user":   adminUser,
			})
		}
	}
}

func AdminAction(c *gin.Context) {
	adminC, exists := c.Get("admin")
	if exists {
		admin := adminC.(database.Admin)

		admin.Name = "JunJun"

		var users []database.User
		if err := gdb.Instance().Model(&database.User{}).Find(&users).Error; err != nil {
			error.SendErrJSON("error", c)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": users,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "no permission",
		"data": gin.H{},
	})
}
