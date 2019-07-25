package routers

import (
	"server/controllers/action"
	"server/controllers/app"
	"server/controllers/article"
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
	cms.Use(jwt.AdminApiJwt(), auth.AdminRequired)
	{
		// real user can do the action
		cms.POST("/users", suaction.AdminAction)
	}

	return r
}
