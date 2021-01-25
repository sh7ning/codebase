package cfg

import (
	"codebase/pkg/app"
	"codebase/pkg/db"
	"codebase/pkg/dingtalk"
	"codebase/pkg/redis"
	"codebase/pkg/web"
)

var Config = &appConf{}

type appConf struct {
	app.Config `mapstructure:",squash"`

	HttpServer *HttpServerConfig `mapstructure:"http_server" validate:"required"`

	DB db.Configs `mapstructure:"db" validate:"required,dive"`

	Redis redis.Configs `mapstructure:"redis" validate:"dive"`

	DingTalk dingtalk.Configs `mapstructure:"dingtalk"`
}

type HttpServerConfig struct {
	web.Config `mapstructure:",squash"`
	Token      string `mapstructure:"token"`
}
