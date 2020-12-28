package user

import (
	"codebase/app/api/app/internal/global"
	"codebase/app/api/app/internal/models"
	"codebase/pkg/log"
	"errors"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

func Create(name string) (*models.User, error) {
	tx := global.DB("")
	model := &models.User{
		Name: name,
	}

	if err := tx.Create(model).Error; err != nil {
		log.Error("user 写入失败", zap.Error(err), zap.String("name", name))
		return nil, err
	}

	return model, nil
}

func Get(userId string) (*models.User, error) {
	tx := global.DB("")
	model := &models.User{}
	q := tx.Where("id = ?", userId).First(model)
	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, q.Error
	}

	return model, nil
}

func List(page, pageSize int) ([]*models.User, int64, error) {
	db := global.DB("")
	var items []*models.User
	var total int64
	err := db.Scopes(models.Paginate(&models.PageParams{
		Page:     page,
		PageSize: pageSize,
	})).Find(&items).Count(&total).Error

	return items, total, err
}
