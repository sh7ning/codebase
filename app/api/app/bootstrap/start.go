package bootstrap

import (
	"codebase/app/api/app/cfg"
	"codebase/app/api/app/services"
	"codebase/app/api/app/web/routes"
	"codebase/pkg/config"
	"codebase/pkg/dingtalk"
	"codebase/pkg/gorm"
	"codebase/pkg/helper"
	"codebase/pkg/log"
	"codebase/pkg/redis"
	"codebase/pkg/web"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"

	"github.com/spf13/pflag"
)

func Start(cfgFile string, flagSet *pflag.FlagSet) {
	//load cfg
	viper, err := config.LoadConfig(cfg.Config, cfgFile, flagSet, nil)
	if err != nil {
		panic(err)
	}

	dingtalks := dingtalk.InitConnections(cfg.Config.DingTalk)

	//init logger
	cfg.Config.InitLoggerConfig()
	log.New(cfg.Config.LoggerConfig)
	defer log.Sync()

	log.Info("using cfg file: " + viper.ConfigFileUsed())
	log.Debug("cfg data", zap.String("config_data", helper.ToJsonString(cfg.Config)))

	dbConnections := gorm.InitConnections(cfg.Config.AppDebug, cfg.Config.DB)
	defer dbConnections.Close()

	redisConnections := redis.InitConnections(cfg.Config.Redis)

	services.InitAppService(dbConnections, redisConnections, dingtalks)

	//运行 api server
	engine := web.NewEngine(cfg.Config.AppDebug)
	routes.Routes(engine)
	httpServer := web.NewServer(engine, cfg.Config.HttpServer)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Error("api server ListenAndServe error, "+err.Error(), zap.Error(err))
		}
	}()

	log.Info("http server run success, listen: " + cfg.Config.HttpServer.Address)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	//todo 增加平滑重启信号相关处理逻辑
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sig := <-quit
	log.Warn("get signal, start shutdown server ...", zap.Any("signal", sig))

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", zap.Error(err))
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Warn("server shutdown with timeout of 2 seconds.")
	}
	log.Warn("Server exiting")
}
