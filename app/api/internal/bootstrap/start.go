package bootstrap

import (
	"codebase/app/api/internal/cfg"
	"codebase/app/api/internal/web"
	"codebase/pkg/db"
	"codebase/pkg/dingtalk"
	"codebase/pkg/log"
	"codebase/pkg/redis"

	"go.uber.org/zap"
)

func Start(errs chan error) error {
	if err := loadResource(); err != nil {
		log.Panic(err.Error(), zap.Error(err))
	}

	//业务代码开始
	web.New()

	return nil
}

func loadResource() error {
	if err := db.Init(cfg.Config.AppDebug, cfg.Config.DB); err != nil {
		return err
	}
	if err := redis.Init(cfg.Config.Redis); err != nil {
		return err
	}
	dingtalk.Init(cfg.Config.DingTalk)

	return nil
}
