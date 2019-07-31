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
	/*
		adminId := c.Query("adminId")

		var admin database.Admin

		val_uint, _ := strconv.ParseUint(adminId, 10, 64)
		var err error
		if admin, err = searchAdminUser(val_uint); err != nil {
			log.Fatalf("req.Auth err: %v", err)
			return
		}

		c.Set("admin", admin)
		c.Next()
	*/
}
