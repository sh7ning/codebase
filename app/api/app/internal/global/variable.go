package global

import (
	"codebase/app/api/app/internal/cfg"
	"codebase/pkg/config"
	"codebase/pkg/db"
	"codebase/pkg/dingtalk"
	"codebase/pkg/helper"
	"codebase/pkg/log"
	"codebase/pkg/redis"

	redisCli "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

var g *Services

type Services struct {
	databases *db.Connections

	redisClients *redis.Connections

	dingTalks *dingtalk.Robots
}

func Init(cfgFile string, flagSet *pflag.FlagSet) {
	//load cfg
	c, err := config.LoadConfig(cfg.Config, cfgFile, flagSet, nil)
	if err != nil {
		panic(err)
	}

	//init logger
	log.New(&log.LoggerConfig{
		Development: cfg.Config.AppDebug,
		AppName:     cfg.Config.AppName,
		LogFile:     cfg.Config.Logger.File,
		Notify:      nil,
	})

	log.Info("using cfg file: " + c.ConfigFileUsed())
	log.Debug("cfg data", zap.String("config_data", helper.ToJsonString(cfg.Config)))

	if g != nil {
		panic("") //todo
	}
	g = &Services{
		databases:    db.New(cfg.Config.AppDebug, cfg.Config.DB),
		redisClients: redis.New(cfg.Config.Redis),
		dingTalks:    dingtalk.New(cfg.Config.DingTalk),
	}
}

func DB(conn string) *gorm.DB {
	return g.databases.Get(conn)
}

func Redis(conn string) *redisCli.Client {
	return g.redisClients.Get(conn)
}
