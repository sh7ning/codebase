package models

import (
	"time"

	"gorm.io/gorm"
)

//CREATE TABLE `users` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
//`name` varchar(255) NOT NULL DEFAULT '',
//`created_at` datetime DEFAULT NULL,
//`updated_at` datetime DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4
type Model struct {
	Id        uint64    `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func Paginate(p *PageParams) func(query *gorm.DB) *gorm.DB {
	return func(query *gorm.DB) *gorm.DB {
		return query.Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize)
	}
}

type PageParams struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"page_size" form:"page_size" binding:"required"`
}

type PageResponse struct {
	Page     int         `json:"current_page"`
	PageSize int         `json:"page_size"`
	Total    int64       `json:"total"`
	Items    interface{} `json:"items"`
}
