package cfg

import (
	"codebase/pkg/config"
	"codebase/pkg/dingtalk"
	"codebase/pkg/gorm"
	"codebase/pkg/redis"
	"codebase/pkg/web"
)

var Config = &AppConf{}

type AppConf struct {
	config.AppConfig `mapstructure:",squash"`

	HttpServer *web.Config `mapstructure:"http_server" validate:"required"`

	DingTalk dingtalk.Configs `mapstructure:"dingtalk"`

	DB gorm.Configs `mapstructure:"db" validate:"required,dive"`

	Redis redis.Configs `mapstructure:"redis" validate:"dive"`
}
