package comment

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
	articleId, _ := c.Get("articleId")

	comments := []database.Comment{}
	if err := gdb.Instance().Find(&comments).Where("articleId = ?", articleId).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, comments)
}

func AddComment(c *gin.Context) {
	articleIdStr := c.Param("id")
	addComment := database.Comment{}
	if err := c.ShouldBindJSON(&addComment); err != nil {
		log.Printf("req.AppInfo err: %v", err)
		return
	}

	articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	addComment.ArticleId = uint(articleId)

	if err := gdb.Instance().Create(&addComment).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, "success")
}

func GetCommentList(c *gin.Context) {
	comments := []database.Comment{}
	if err := gdb.Instance().Preload("User").Find(&comments).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"comments": comments,
	})
}

// For Admin Api
func PublishComment(c *gin.Context) {
	commentIdStr := c.Param("id")
	publishBeforeComment := database.Comment{}
	// 削除したいレコードのIDを指定
	commentId, err := strconv.ParseUint(commentIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		error.SendErrJSON("error", c)
		return
	}
	publishBeforeComment.ID = uint(commentId)

	if err := gdb.Instance().Find(&publishBeforeComment).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	publishAfterComment := publishBeforeComment
	publishAfterComment.Status = 2
	if err := gdb.Instance().Model(&publishBeforeComment).Update(&publishAfterComment).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Save(&publishAfterComment).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"err_msg": "success",
	})
}

func DeleteComment(c *gin.Context) {
	commentIdStr := c.Param("id")
	delComment := database.Comment{}
	// 削除したいレコードのIDを指定
	commentId, err := strconv.ParseUint(commentIdStr, 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	delComment.ID = uint(commentId)

	if err := gdb.Instance().Find(&delComment).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	if err := gdb.Instance().Delete(&delComment).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "",
	})
}
