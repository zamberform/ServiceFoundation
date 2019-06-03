package action

import (
	"server/model"
	"server/pkg/gdb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DoAction(c *gin.Context) {
	userC, exists := c.Get("user")
	var user model.User
	if exists {
		user = userC.(model.User)

		user.Introduce += "admin"
		if err := gdb.Instance().Model(&user).Update("introduce", user.Introduce); err != nil {

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": gin.H{},
	})
}
