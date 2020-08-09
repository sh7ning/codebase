package routes

import (
	"codebase/app/api/app/web/api/home"
	"codebase/app/api/app/web/api/user"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	engine.GET("/index", home.Home())

	apiRoutes := engine.Group("/v1")
	{
		//获取用户信息
		apiRoutes.GET("/user/:userId", user.Get)
		//创建用户
		apiRoutes.GET("/user", user.Create)
	}
}