package user

import (
	"codebase/app/api/app/internal/models"
	"codebase/app/api/app/internal/repositories/user"
	"codebase/app/api/app/internal/services/user/params"
	"codebase/pkg/log"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func Create(tx *gorm.DB, form params.UserCreateRequest) (*models.User, error) {
	model, err := user.Create(tx, form.Name)

	if err != nil {
		log.Error("user 写入失败", zap.Error(err), zap.String("name", form.Name))

		return nil, nil
	}

	return model, nil
}

func Get(tx *gorm.DB, userId string) (*models.User, error) {
	return user.Get(tx, userId)
}
