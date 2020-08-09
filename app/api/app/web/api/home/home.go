package home

import (
	"codebase/pkg/web/response"
	"time"

	"github.com/gin-gonic/gin"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.SuccessJSON(c, time.Now().Unix())
	}
}
