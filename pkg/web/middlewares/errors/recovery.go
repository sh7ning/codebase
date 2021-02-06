package errors

import (
	"codebase/pkg/log"
	"codebase/pkg/web/response"
	"fmt"
	"net/http/httputil"
	"runtime"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				_, file, line, _ := runtime.Caller(2)
				request, _ := httputil.DumpRequest(c.Request, false)
				log.Error(
					"[Recovery] panic recovered",
					zap.String("line", fmt.Sprintf("%s:%d", file, line)),
					zap.Any("err", err),
					zap.String("request", string(request)),
				)
				response.ErrorJson(c, response.CodeSystemError, "")
			}
		}()
		c.Next()
	}

	//return gin.Recovery()
}
