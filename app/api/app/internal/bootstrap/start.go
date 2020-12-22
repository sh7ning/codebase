package bootstrap

import (
	"codebase/app/api/app/internal/cfg"
	"codebase/app/api/app/internal/services"
	"codebase/app/api/app/internal/web/routes"
	"codebase/pkg/config"
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

	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

func Start(cfgFile string, flagSet *pflag.FlagSet) {
	//load cfg
	c, err := config.LoadConfig(cfg.Config, cfgFile, flagSet, nil)
	if err != nil {
		panic(err)
	}

	//init logger
	cfg.Config.InitLoggerConfig()
	log.New(cfg.Config.LoggerConfig)
	defer log.Sync()

	log.Info("using cfg file: " + c.ConfigFileUsed())
	log.Debug("cfg data", zap.String("config_data", helper.ToJsonString(cfg.Config)))

	dbConnections := gorm.InitConnections(cfg.Config.AppDebug, cfg.Config.DB)
	defer dbConnections.Close()

	redisConnections := redis.InitConnections(cfg.Config.Redis)

	services.InitAppService(dbConnections, redisConnections, cfg.Config.DingTalk)

	//运行 api server
	engine := web.NewEngine(cfg.Config.AppDebug, cfg.Config.HttpServer)
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", zap.Error(err))
	}

	log.Info("Server exiting")
}
