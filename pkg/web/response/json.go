package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func SuccessJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Response{
		Code: CodeSuccess,
		Data: data,
		Msg:  "",
	})
}

func ErrorJson(c *gin.Context, code int, msg string) {
	ErrorJsonWithStatusCode(c, http.StatusOK, code, nil, msg)
}

func ErrorJsonWithData(c *gin.Context, code int, data interface{}, msg string) {
	ErrorJsonWithStatusCode(c, http.StatusOK, code, data, msg)
}

func ErrorJsonWithStatusCode(c *gin.Context, statusCode, code int, data interface{}, msg string) {
	c.JSON(statusCode, &Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
