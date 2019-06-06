package suaction

import (
	"net/http"
	"server/models/db"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func AdminAction(c *gin.Context) {
	userC, exists := c.Get("user")
	var user db.User
	if exists {
		user = userC.(db.User)

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
