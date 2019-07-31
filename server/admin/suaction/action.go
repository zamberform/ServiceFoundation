package suaction

import (
	"log"
	"net/http"
	"server/controllers/error"
	"server/models/database"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	if err := bcrypt.CompareHashAndPassword([]byte(adminUser.Pass), []byte(currentUser.Password)); err != nil {
		log.Printf("get.db.Signin err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})
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
