package app

import (
	"api/model"
	"api/pkg/gdb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppInfo(c *gin.Context) {
	var appInfo model.App
	platformId, _ := c.Get("platform")
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
