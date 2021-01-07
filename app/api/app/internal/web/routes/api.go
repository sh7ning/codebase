package routes

import (
	"codebase/app/api/app/internal/web/controllers/home"
	"codebase/app/api/app/internal/web/controllers/user"

	"github.com/gin-gonic/gin"
)

func Routes(engine *gin.Engine) {
	engine.GET("/", home.Home())

	apiRoutes := engine.Group("/v1")
	{
		//获取用户列表
		apiRoutes.GET("/users", user.List)
		//获取用户信息
		apiRoutes.GET("/user/:userId", user.Get)
		//创建用户
		apiRoutes.GET("/user", user.Create)
	}
}
