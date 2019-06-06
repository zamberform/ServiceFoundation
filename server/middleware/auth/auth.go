package auth

import (
	"server/models/db"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func SigninRequired(c *gin.Context) {
	var user db.User
	var err error
	if user, err = searchUser(c); err != nil {

		return
	}

	if user.Status <= 0 {
		return
	}

	c.Set("user", user)
	c.Next()
}

func AdminRequired(c *gin.Context) {
	var user db.User
	var err error
	if user, err = searchUser(c); err != nil {

		return
	}

	if user.Status <= 0 {
		return
	}

	c.Set("user", user)
	c.Next()
}

func VipReqired(c *gin.Context) {
	userC, exists := c.Get("user")
	var user db.User
	if exists {
		user = userC.(db.User)

		user.Introduce += "vip"
		if err := gdb.Instance().Model(&user).Update("introduce", user.Introduce); err != nil {

			return
		}

		c.Set("user", user)
		c.Next()
	}
}
