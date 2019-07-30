package article

import (
	"fmt"
	"log"
	"net/http"
	"server/controllers/error"
	"server/models/database"
	"server/pkg/gdb"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	articles := []database.Article{}
	if err := gdb.Instance().Find(&articles).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, articles)
}

// For Admin Api
func AddArticle(c *gin.Context) {
	addArticle := database.Article{}
	if err := c.ShouldBindJSON(&addArticle); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}

	if err := gdb.Instance().Create(&addArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, "success")
}

func UpdateArticle(c *gin.Context) {
	articleIdStr := c.Param("id")
	updateInfo := database.Article{}
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		log.Fatalf("req.AppInfo err: %v", err)
		return
	}

	publishBeforeArticle := database.Article{}
	// 削除したいレコードのIDを指定
	articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	publishBeforeArticle.ID = uint(articleId)

	if err := gdb.Instance().Find(&publishBeforeArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	publishAfterArticle := publishBeforeArticle
	publishAfterArticle.Title = updateInfo.Title
	publishAfterArticle.ContentDesc = updateInfo.ContentDesc

	if err := gdb.Instance().Model(&publishBeforeArticle).Update(&publishAfterArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	if err := gdb.Instance().Save(&publishAfterArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	c.JSON(http.StatusOK, "success")
}

func PublishArticle(c *gin.Context) {
	articleIdStr := c.Param("id")
	publishBeforeArticle := database.Article{}
	// 削除したいレコードのIDを指定
	articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	publishBeforeArticle.ID = uint(articleId)

	if err := gdb.Instance().Find(&publishBeforeArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	publishAfterArticle := publishBeforeArticle
	publishAfterArticle.Status = 3
	if err := gdb.Instance().Model(&publishBeforeArticle).Update(&publishAfterArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	if err := gdb.Instance().Save(&publishAfterArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	c.JSON(http.StatusOK, "success")
}

func HideArticle(c *gin.Context) {
	articleIdStr := c.Param("id")
	publishBeforeArticle := database.Article{}
	// 削除したいレコードのIDを指定
	articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	publishBeforeArticle.ID = uint(articleId)

	if err := gdb.Instance().Find(&publishBeforeArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	publishAfterArticle := publishBeforeArticle
	publishAfterArticle.Status = 2
	if err := gdb.Instance().Model(&publishBeforeArticle).Update(&publishAfterArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	if err := gdb.Instance().Save(&publishAfterArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	c.JSON(http.StatusOK, "success")
}

func DeleteArticle(c *gin.Context) {
	articleIdStr := c.Param("id")
	delArticle := database.Article{}
	// 削除したいレコードのIDを指定
	articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	delArticle.ID = uint(articleId)

	if err := gdb.Instance().Find(&delArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Delete(&delArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
	}

	c.JSON(http.StatusOK, "success")
}
