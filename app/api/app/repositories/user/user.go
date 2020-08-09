package user

import (
	"codebase/app/api/app/models"

	"github.com/jinzhu/gorm"
)

func Create(tx *gorm.DB, name string) (*models.User, error) {
	model := models.User{
		Name: name,
	}

	if err := tx.Create(&model).Error; err != nil {
		return nil, err
	}

	return &model, nil
}

func Get(tx *gorm.DB, userId string) (*models.User, error) {
	model := models.User{}
	q := tx.Where("id = ?", userId).First(&model)
	if q.Error != nil {
		if q.RecordNotFound() {
			return nil, nil
		}

		return nil, q.Error
	}

	return &model, nil
}
