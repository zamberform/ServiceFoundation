package article

import (
	"fmt"
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
	articles := []database.Article{}
	if err := gdb.Instance().Find(&articles).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "",
		"count":    len(articles),
		"articles": articles,
	})
}

// For Admin Api
func AddArticle(c *gin.Context) {
	addArticle := database.Article{}
	if err := c.ShouldBindJSON(&addArticle); err != nil {
		log.Fatalf("add article info err: %v", err)
		return
	}

	checkDefinedArticle := database.Article{}
	if err := gdb.Instance().Find(&checkDefinedArticle).Where("title = ?", addArticle.Title).Error; err != nil {
		log.Printf("重複している: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	addArticle.CreatedAt = time.Now()
	addArticle.UpdatedAt = time.Now()
	if err := gdb.Instance().Create(&addArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})
}

func UpdateArticle(c *gin.Context) {
	updateInfo := database.Article{}
	if err := c.ShouldBindJSON(&updateInfo); err != nil {
		log.Printf("req.AppInfo err: %v", err)
		return
	}

	publishBeforeArticle := database.Article{}

	if err := gdb.Instance().Find(&publishBeforeArticle).Where("id = ?", updateInfo.ID).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	publishAfterArticle := publishBeforeArticle
	publishAfterArticle.Title = updateInfo.Title
	publishAfterArticle.ContentDesc = updateInfo.ContentDesc
	publishAfterArticle.UpdatedAt = time.Now()
	publishAfterArticle.CommentFlg = updateInfo.CommentFlg
	if err := gdb.Instance().Model(&publishBeforeArticle).Update(&publishAfterArticle).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Save(&publishAfterArticle).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})
}

func PublishArticle(c *gin.Context) {
	articleIdStr := c.Param("id")
	publishBeforeArticle := database.Article{}
	// 削除したいレコードのIDを指定
	articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		return
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
		return
	}

	if err := gdb.Instance().Save(&publishAfterArticle).Error; err != nil {
		log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
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
		error.SendErrJSON("error", c)
		return
	}
	publishBeforeArticle.ID = uint(articleId)

	if err := gdb.Instance().Find(&publishBeforeArticle).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	publishAfterArticle := publishBeforeArticle
	publishAfterArticle.Status = 2
	if err := gdb.Instance().Model(&publishBeforeArticle).Update(&publishAfterArticle).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Save(&publishAfterArticle).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"err_msg": "success",
	})
}

func FindArticle(c *gin.Context) {
	articleIdStr := c.Param("id")
	findArticle := database.Article{}
	// 削除したいレコードのIDを指定
	articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	findArticle.ID = uint(articleId)

	if err := gdb.Instance().Find(&findArticle).Error; err != nil {
		// log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"msg":     "success",
		"article": findArticle,
	})
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

	if err := gdb.Instance().First(&delArticle).Error; err != nil {
		// log.Fatalf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Delete(&delArticle).Error; err != nil {
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"err_msg": "success",
	})
}
