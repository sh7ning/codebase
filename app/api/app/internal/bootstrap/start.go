package bootstrap

import (
	"codebase/app/api/app/internal/cfg"
	"codebase/app/api/app/internal/web"
	"codebase/pkg/db"
	"codebase/pkg/dingtalk"
	"codebase/pkg/redis"
)

func Start() error {
	loadResource()

	//业务代码开始
	web.New()

	return nil
}

func loadResource() {
	db.Init(cfg.Config.AppDebug, cfg.Config.DB)
	redis.Init(cfg.Config.Redis)
	dingtalk.Init(cfg.Config.DingTalk)
}
