package cfg

import (
	"codebase/pkg/config"
	"codebase/pkg/db"
	"codebase/pkg/dingtalk"
	"codebase/pkg/redis"
	"codebase/pkg/web"
)

var Config = &appConf{}

type appConf struct {
	config.AppConfig `mapstructure:",squash"`

	HttpServer *web.Config `mapstructure:"http_server" validate:"required"`

	DB db.Configs `mapstructure:"db" validate:"required,dive"`

	Redis redis.Configs `mapstructure:"redis" validate:"dive"`

	DingTalk dingtalk.Configs `mapstructure:"dingtalk"`
}
