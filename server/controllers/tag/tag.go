package tag

import (
	"log"
	"net/http"
	"server/controllers/error"
	"server/models/database"
	"server/pkg/gdb"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	tags := []database.Tag{}
	if err := gdb.Instance().Find(&tags).Error; err != nil {
		log.Fatalf("not tags err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
		"tags":   tags,
	})
}

// For Admin Api
func AddTag(c *gin.Context) {
	addTag := database.Tag{}
	if err := c.ShouldBindJSON(&addTag); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}

	addTag.CreatedAt = time.Now()
	addTag.UpdatedAt = time.Now()
	if err := gdb.Instance().Create(&addTag).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
	})
}

func UpdateTag(c *gin.Context) {
	updateInfo := database.Tag{}
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}

	updateBeforeTag := database.Tag{}
	updateBeforeTag.ID = updateInfo.ID

	if err := gdb.Instance().Find(&updateBeforeTag).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	updateAfterTag := updateBeforeTag
	updateAfterTag.Name = updateInfo.Name
	updateAfterTag.Color = updateInfo.Color
	updateAfterTag.UpdatedAt = time.Now()
	if err := gdb.Instance().Model(&updateBeforeTag).Update(&updateAfterTag).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	if err := gdb.Instance().Save(&updateAfterTag).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
	})
}

func DeleteTag(c *gin.Context) {
	tagIdStr := c.Param("id")
	delTag := database.Tag{}
	// 削除したいレコードのIDを指定
	tagId, err := strconv.ParseUint(tagIdStr, 10, 32)
	if err != nil {
		log.Printf("Delete Tag Err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	delTag.ID = uint(tagId)

	if err := gdb.Instance().Find(&delTag).Error; err != nil {
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Delete(&delTag).Error; err != nil {
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
	})
}
