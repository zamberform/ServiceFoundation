package routers

import (
	"server/controllers/action"
	"server/controllers/app"
	"server/controllers/user"

	"server/admin/suaction"

	"server/middleware/auth"
	"server/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func InitRouter(apiPrefix string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// check app version
	r.POST("/app", app.AppInfo)
	// had signin next will login
	r.POST("/login", user.Login)
	apis := r.Group(apiPrefix)
	apis.Use(jwt.ApiJwt())
	{
		// change to real user
		apis.POST("/signin", user.SignIn)
		// real user can do the action
		apis.POST("/actions", auth.SigninRequired, action.DoAction)
		// vip user can do
		apis.POST("/vip", auth.SigninRequired, auth.VipReqired, action.DoAction)
		// real user can quit
		apis.POST("/quits", user.Withdrawal)
	}

	cms := r.Group("/cms", auth.AdminRequired)
	cms.Use(jwt.ApiJwt())
	{
		// real user can do the action
		cms.POST("/users", suaction.AdminAction)
	}

	return r
}
