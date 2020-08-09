package models

import "time"

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
