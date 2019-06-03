package auth

import (
	"api/model"

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

func VipReqired(c *gin.Context) {
	var user model.User
	var err error
	if user, err = searchUser(c); err != nil {

		return
	}
	if user.Status <= 0 {
		return
	}

	if user.Role <= 3 {
		return
	}
	c.Set("user", user)
	c.Next()
}
