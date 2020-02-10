package auth

import (
	"app/pkg/cfg"
	"app/pkg/web/response"

	"github.com/gin-gonic/gin"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token := c.GetHeader("X-API-TOKEN"); token != cfg.AppConfig.HttpServer.Token {
			response.ErrorJson(c, response.CodeUnauthorized, "token error")
			c.Abort()
			return
		}
		c.Next()
	}
}
