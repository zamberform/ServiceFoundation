package action

import (
	"log"
	"net/http"
	"server/models/db"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func DoAction(c *gin.Context) {
	userC, exists := c.Get("user")
	var user db.User
	if exists {
		user = userC.(db.User)

		user.Introduce += "admin"
		if err := gdb.Instance().Model(&user).Update("introduce", user.Introduce); err != nil {
			log.Fatalf("get.db.AppInfo err: %v", err)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": gin.H{},
	})
}
