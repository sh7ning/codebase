package services

import (
	"codebase/pkg/dingtalk"
	"codebase/pkg/gorm"
	"codebase/pkg/redis"
)

var AppService *Service

type Service struct {
	DbConnections *gorm.Connections

	RedisConnections *redis.Connections

	Dingtalks *dingtalk.Connections
}

func InitAppService(
	dbConnections *gorm.Connections,
	redisConnections *redis.Connections,
	dingtalks *dingtalk.Connections) {
	AppService = &Service{
		DbConnections:    dbConnections,
		RedisConnections: redisConnections,
		Dingtalks:        dingtalks,
	}
}
