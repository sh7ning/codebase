package user

import (
	"app/pkg/models"
	"app/pkg/repositories/user"
	"app/pkg/services/log"
	"app/pkg/services/user/params"

	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
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
