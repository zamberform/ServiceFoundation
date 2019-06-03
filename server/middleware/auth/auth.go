package auth

import (
	"server/model"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func SigninRequired(c *gin.Context) {
	var user model.User
	var err error
	if user, err = searchUser(c); err != nil {

		return
	}

	if user.Status <= 0 {
		return
	}

	if user.Role <= 1 {
		return
	}
	c.Set("user", user)
	c.Next()
}

func AdminRequired(c *gin.Context) {
	var user model.User
	var err error
	if user, err = searchUser(c); err != nil {

		return
	}

	if user.Status <= 0 {
		return
	}

	if user.Role <= 1 {
		return
	}
	c.Set("user", user)
	c.Next()
}

func VipReqired(c *gin.Context) {
	userC, exists := c.Get("user")
	var user model.User
	if exists {
		user = userC.(model.User)

		user.Introduce += "vip"
		if err := gdb.Instance().Model(&user).Update("introduce", user.Introduce); err != nil {

			return
		}

		c.Set("user", user)
		c.Next()
	}
}
