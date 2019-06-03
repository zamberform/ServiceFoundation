package routers

import (
	"api/controllers/action"
	"api/controllers/app"
	"api/controllers/user"
	"api/middleware/auth"
	"api/middleware/jwt"

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

	return r
}
