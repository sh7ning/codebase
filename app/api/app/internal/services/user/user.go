package user

import (
	"codebase/app/api/app/internal/global"
	"codebase/app/api/app/internal/models"
	"codebase/app/api/app/internal/repositories/user"
	"codebase/pkg/log"

	"go.uber.org/zap"
)

func Create(name string) (*models.User, error) {
	tx := global.DB("")
	model, err := user.Create(tx, name)

	if err != nil {
		log.Error("user 写入失败", zap.Error(err), zap.String("name", name))

		return nil, nil
	}

	return model, nil
}

func Get(userId string) (*models.User, error) {
	tx := global.DB("")
	return user.Get(tx, userId)
}
