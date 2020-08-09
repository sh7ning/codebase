package user

import (
	"codebase/app/api/app/services"
	"codebase/app/api/app/services/user"
	"codebase/app/api/app/services/user/params"
	"codebase/pkg/log"
	"codebase/pkg/web/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

func Get(c *gin.Context) {
	requestId := c.Param("userId")
	tx := services.AppService.DbConnections.Connection("")
	model, err := user.Get(tx, requestId)
	if err != nil {
		log.Error("获取用户失败, userId: "+requestId+", error: "+err.Error(), zap.Error(err))
		response.ErrorJson(c, response.CodeSystemError, err.Error())
		return
	}

	if model == nil {
		response.ErrorJson(c, response.CodeUserNotFound, "user not found")
		return
	}

	response.SuccessJSON(c, gin.H{
		"id":   model.Id,
		"name": model.Name,
	})
}

func Create(c *gin.Context) {
	var form params.UserCreateRequest
	if err := c.ShouldBindWith(&form, binding.Form); err != nil {
		log.Debug("Create CodeParamsError", zap.Error(err))
		//response.ErrorJson(c, response.CodeParamsError, "")
		response.ErrorJson(c, response.CodeParamsError, err.Error())
		return
	}

	tx := services.AppService.DbConnections.Connection("")
	model, err := user.Create(tx, form)

	if err != nil {
		response.ErrorJson(c, response.CodeSystemError, err.Error())
		return
	}

	response.SuccessJSON(c, gin.H{
		"id":   model.Id,
		"name": model.Name,
	})
}