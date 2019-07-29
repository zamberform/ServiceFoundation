package routers

import (
	"server/controllers/action"
	"server/controllers/app"
	"server/controllers/article"
	"server/controllers/comment"
	"server/controllers/tag"
	"server/controllers/user"

	"server/admin/suaction"

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
		apis.POST("/comment/list/:articleId", auth.SigninRequired, comment.GetAll)
		apis.POST("/comment/:articleId", auth.SigninRequired, auth.VipReqired, comment.AddComment)
		apis.POST("/signin", user.SignIn)
		// real user can do the action
		apis.POST("/actions", auth.SigninRequired, action.DoAction)
		// vip user can do
		apis.POST("/vip", auth.SigninRequired, auth.VipReqired, action.DoVipAction)
		// real user can quit
		apis.POST("/quits", user.Withdrawal)
	}

	cms := r.Group(cmsPrefix)
	cms.Use(jwt.AdminApiJwt(), auth.AdminRequired)
	{
		// real user can do the action
		cms.POST("/users", suaction.AdminAction)
		cms.POST("/tags", suaction.AdminAction)
		cms.POST("/comments", suaction.AdminAction)
		cms.POST("/articles", suaction.AdminAction)

		cms.DELETE("/article/:id", suaction.AdminAction)
		cms.DELETE("/tag/:id", suaction.AdminAction)
		cms.DELETE("/user/:id", suaction.AdminAction)
		cms.DELETE("/comment/:id", suaction.AdminAction)

		cms.POST("/article/add", suaction.AdminAction)
		cms.POST("/tag/add", suaction.AdminAction)
		cms.POST("/user/add", suaction.AdminAction)
		cms.POST("/comment/add", suaction.AdminAction)

		cms.POST("/article/:id", suaction.AdminAction)
		cms.POST("/tag/:id", suaction.AdminAction)
		cms.POST("/user/:id", suaction.AdminAction)
		cms.POST("/comment/:id", suaction.AdminAction)
	}

	return r
}
