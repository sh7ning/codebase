package user

import (
	"codebase/app/api/internal/models"
	"codebase/pkg/db"
	"codebase/pkg/log"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Create(name string) (*models.User, error) {
	conn := db.Conn("")
	model := &models.User{
		Name: name,
	}

	if err := conn.Create(model).Error; err != nil {
		log.Error("user 写入失败", zap.Error(err), zap.String("name", name))
		return nil, err
	}

	return model, nil
}

func List(page, pageSize int) ([]*models.User, int64, error) {
	conn := db.Conn("")
	var items []*models.User
	var total int64
	err := conn.Scopes(models.Paginate(&models.PageParams{
		Page:     page,
		PageSize: pageSize,
	})).Find(&items).Count(&total).Error

	return items, total, err
}

func Get(userId string) (*models.User, error) {
	conn := db.Conn("")
	model := &models.User{}
	q := conn.Where("id = ?", userId).First(model)
	if q.Error != nil {
		if errors.Is(q.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, q.Error
	}

	return model, nil
}
