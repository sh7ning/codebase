package user

import (
	"codebase/app/api/app/internal/models"
	"codebase/app/api/app/internal/services/user"
	"codebase/app/api/app/internal/web/controllers/user/params"
	"codebase/pkg/helper"
	"codebase/pkg/log"
	"codebase/pkg/web/response"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	var form params.UserCreateRequest
	if err := c.ShouldBindWith(&form, binding.Form); err != nil {
		log.Debug("Create CodeParamsError", zap.Error(err))
		//response.ErrorJson(c, response.CodeParamsError, "")
		response.ErrorJson(c, response.CodeParamsError, err.Error())
		return
	}

	model, err := user.Create(form.Name)

	if err != nil {
		response.ErrorJson(c, response.CodeSystemError, err.Error())
		return
	}

	response.SuccessJSON(c, gin.H{
		"id":   model.Id,
		"name": model.Name,
	})
}

func List(c *gin.Context) {
	var form params.UserListRequest
	if err := c.ShouldBindWith(&form, binding.Form); err != nil {
		log.Debug("List CodeParamsError", zap.Error(err))
		response.ErrorJson(c, response.CodeParamsError, err.Error())
		return
	}

	data, total, err := user.List(form.Page, form.PageSize)
	if err != nil {
		log.Error("获取用户列表失败, request: "+helper.ToJsonString(form), zap.Error(err))
		response.ErrorJson(c, response.CodeSystemError, err.Error())
		return
	}

	response.SuccessJSON(c, models.PageResponse{
		Page:     form.Page,
		PageSize: form.PageSize,
		Total:    total,
		Items:    data,
	})
}

func Get(c *gin.Context) {
	requestId := c.Param("userId")
	model, err := user.Get(requestId)
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
