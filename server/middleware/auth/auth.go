package auth

import (
	"log"
	"server/models/database"
	"server/models/request"

	"github.com/gin-gonic/gin"
)

func SigninRequired(c *gin.Context) {
	var commonReq request.CommonReq
	commonInfo, _ := c.Get("common")

	commonReq = commonInfo.(request.CommonReq)
	userId := commonReq.User.UserId

	var user database.User
	var err error
	if user, err = searchUser(userId); err != nil {
		log.Fatalf("req.Auth err: %v", err)
		return
	}

	c.Set("user", user)
	c.Next()
}

func VipReqired(c *gin.Context) {
	var commonReq request.CommonReq
	commonInfo, _ := c.Get("common")

	commonReq = commonInfo.(request.CommonReq)
	userId := commonReq.User.UserId

	var user database.User
	var err error
	if user, err = searchVipUser(userId); err != nil {
		log.Fatalf("req.Auth err: %v", err)
		return
	}

	c.Set("user", user)
	c.Next()
}

func AdminRequired(c *gin.Context) {
	c.Next()
}
