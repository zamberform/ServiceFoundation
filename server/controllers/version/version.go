package version

import (
	"log"
	"net/http"
	"server/controllers/error"
	"server/models/database"
	"server/pkg/gdb"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	versions := []database.App{}
	if err := gdb.Instance().Find(&versions).Error; err != nil {
		log.Printf("get.db.App err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"versions": versions,
	})
}
