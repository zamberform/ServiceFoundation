package suaction

import (
	"net/http"
	"server/models/database"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func AdminAction(c *gin.Context) {
	userC, exists := c.Get("user")
	var user database.User
	if exists {
		user = userC.(database.User)

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
