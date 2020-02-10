package home

import (
	"time"

	"app/pkg/web/response"

	"github.com/gin-gonic/gin"
)

func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.SuccessJSON(c, time.Now().Unix())
	}
}
