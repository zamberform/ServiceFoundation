package suaction

import (
	"net/http"
	"server/models/database"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func AdminAction(c *gin.Context) {
	adminC, exists := c.Get("admin")
	if exists {
		admin := adminC.(database.Admin)

		admin.Name = "JunJun"

		var users []database.User
		if err := gdb.Instance().Model(&database.User{}).Find(&users).Error; err != nil {

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
