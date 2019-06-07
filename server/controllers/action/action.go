package action

import (
	"log"
	"net/http"
	"server/models/database"
	"server/models/response"
	"server/pkg/codes"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func DoAction(c *gin.Context) {
	var reqUser database.User
	userInfo, _ := c.Get("user")

	reqUser = userInfo.(database.User)

	reqUser.Introduce += "do_action"
	if err := gdb.Instance().Model(&database.User{}).Where("id = ?", reqUser.ID).Updates(reqUser).Error; err != nil {
		log.Fatalf("get.db.Signin err: %v", err)
		return
	}

	var response response.CommonRes
	response.Msg = "success"
	response.Code = codes.SUCCESS

	c.JSON(http.StatusOK, response)
}

func DoVipAction(c *gin.Context) {
	var reqUser database.User
	userInfo, _ := c.Get("user")

	reqUser = userInfo.(database.User)

	reqUser.Introduce += "Vip"
	if err := gdb.Instance().Model(&database.User{}).Where("id = ?", reqUser.ID).Updates(reqUser).Error; err != nil {
		log.Fatalf("get.db.Signin err: %v", err)
		return
	}

	var response response.CommonRes
	response.Msg = "success"
	response.Code = codes.SUCCESS

	c.JSON(http.StatusOK, response)
}
