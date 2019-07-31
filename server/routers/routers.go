package routers

import (
	"server/admin/suaction"
	"server/controllers/action"
	"server/controllers/app"
	"server/controllers/article"
	"server/controllers/comment"
	"server/controllers/tag"
	"server/controllers/user"

	"server/middleware/auth"
	"server/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter(apiPrefix string, cmsPrefix string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/app", app.AppInfo)
	apis := r.Group(apiPrefix)
	apis.Use(jwt.ApiJwt())
	{
		apis.POST("/article/list", article.GetAll)
		apis.POST("/tag/list", auth.SigninRequired, tag.GetAll)
		apis.POST("/comment/list", auth.SigninRequired, comment.GetAll)
		apis.POST("/comment", auth.SigninRequired, auth.VipReqired, comment.AddComment)
		apis.POST("/signin", user.SignIn)
		// real user can do the action
		apis.POST("/actions", auth.SigninRequired, action.DoAction)
		// vip user can do
		apis.POST("/vip", auth.SigninRequired, auth.VipReqired, action.DoVipAction)
		// real user can quit
		apis.POST("/quits", user.Withdrawal)
	}

	cms := r.Group(cmsPrefix)
	cms.POST("/login", suaction.Login)
	cms.Use(jwt.AdminApiJwt(), auth.AdminRequired)
	{
		// real user can do the action
		cms.POST("/users", user.GetUserList)
		cms.POST("/tags", tag.GetAll)
		cms.POST("/comments", comment.GetCommentList)
		cms.POST("/articles", article.GetAll)

		cms.DELETE("/article/:id", article.DeleteArticle)
		cms.DELETE("/tag/:id", tag.DeleteTag)
		cms.DELETE("/user/:id", user.DeleteUser)
		cms.DELETE("/comment/:id", comment.DeleteComment)

		cms.POST("/article/add", article.AddArticle)
		cms.POST("/tag/add", tag.AddTag)
		cms.POST("/user/add", user.AddUser)
		cms.POST("/comment/add", comment.AddComment)

		cms.POST("/article", article.UpdateArticle)
		cms.POST("/tag", tag.UpdateTag)
		cms.POST("/user", user.UpdateUser)
		cms.POST("/comment", comment.PublishComment)

		cms.POST("/article/push", article.PublishArticle)
	}

	return r
}
