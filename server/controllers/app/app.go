package app

import (
	"log"
	"net/http"
	"server/controllers/error"
	"server/models/db"
	"server/models/request"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func AppInfo(c *gin.Context) {
	var commonRequest request.CommonReq
	var appInfo db.App
	if err := c.ShouldBindJSON(&commonRequest); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	platformId := commonRequest.Platform
	if err := gdb.Instance().Where("platform_type = ?", platformId).Find(&appInfo).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"name":         appInfo.Name,
			"version":      appInfo.Version,
			"platform":     appInfo.PlatformType,
			"updateStatus": appInfo.UpdateStatus,
			"updateURL":    appInfo.URL,
		},
	})
}
