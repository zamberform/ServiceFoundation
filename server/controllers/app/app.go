package app

import (
	"net/http"
	"server/controllers/models"
	"server/model"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func AppInfo(c *gin.Context) {
	var commonRequest models.CommonRequest
	var appInfo model.App
	if err := c.ShouldBindJSON(&commonRequest); err != nil {
		return
	}
	platformId := commonRequest.Platform
	if err := gdb.Instance().First(&appInfo, platformId).Error; err != nil {

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"name":         appInfo.Name,
			"version":      appInfo.Version,
			"platform":     appInfo.AppPlatform,
			"updateStatus": appInfo.UpdateStatus,
			"updateURL":    appInfo.AppURL,
		},
	})
}
