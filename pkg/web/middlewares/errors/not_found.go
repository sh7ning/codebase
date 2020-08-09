package errors

import (
	"codebase/pkg/web/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.ErrorJsonWithStatusCode(c, http.StatusNotFound, response.CodeNotFound, "api not found")
		//http.NotFound(c.Writer, c.Request)
	}
}
