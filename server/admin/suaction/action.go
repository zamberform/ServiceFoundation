package suaction

import (
	"net/http"
	"server/model"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func AdminAction(c *gin.Context) {
	userC, exists := c.Get("user")
	var user model.User
	if exists {
		user = userC.(model.User)

		user.Introduce += "action"
		if err := gdb.Instance().Model(&user).Update("introduce", user.Introduce); err != nil {

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": gin.H{},
	})
}