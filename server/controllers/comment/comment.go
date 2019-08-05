package comment

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

type Comment struct {
	User      string `form:"username" json:"username" xml:"username"`
	Email     string `form:"email" json:"email" xml:"email"`
	Content   string `form:"content" json:"content" xml:"content"`
	ArticleId string `form:"articleId" json:"articleId" xml:"articleId"`
}

func GetAll(c *gin.Context) {
	articleId, _ := c.Get("articleId")

	comments := []database.Comment{}
	if err := gdb.Instance().Find(&comments).Where("articleId = ?", articleId).Error; err != nil {
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

func GetArticleComments(c *gin.Context) {
	articleId, _ := c.Get("id")

	comments := []database.Comment{}
	if err := gdb.Instance().Preload("User").Find(&comments).Where("articleId = ?", articleId).Error; err != nil {
		log.Printf("get.db.AppInfo err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"msg":      "success",
		"count":    len(comments),
		"comments": comments,
	})
}

func AddComment(c *gin.Context) {
	addComment := Comment{}
	if err := c.ShouldBindJSON(&addComment); err != nil {
		log.Printf("add comment err: %v", err)
		return
	}

	articleId, err := strconv.ParseUint(addComment.ArticleId, 10, 32)
	if err != nil {
		fmt.Println(err)
		error.SendErrJSON("error", c)
		return
	}

	dbComment := database.Comment{}
	dbComment.CreatedAt = time.Now()
	dbComment.Comment = addComment.Content
	var user database.User
	if err := gdb.Instance().Where("name = ?", addComment.User).Find(&user).Error; err == nil {
		dbComment.UserId = user.ID
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 1111,
			"msg":    "先にログインしてください",
		})
		return
	}
	var article database.Article
	if err := gdb.Instance().Where("id = ?", uint(articleId)).Find(&article).Error; err == nil {
		dbComment.ArticleId = article.ID
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 1111,
			"msg":    "対象ブログは見つからない",
		})
		return
	}

	if err := gdb.Instance().Preload("User").Create(&dbComment).Error; err != nil {
		log.Printf("add comment err: %v", err)
		error.SendErrJSON("error", c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "success",
	})
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
