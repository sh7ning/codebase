package bootstrap

import (
	"app/pkg/cfg"
	"app/pkg/services/db"
	"app/pkg/services/dingtalk"
	"app/pkg/services/log"
	"app/pkg/services/redis"
	"app/pkg/utils/helper"
	"app/pkg/web"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

func Run(cfgFile string, flagSet *pflag.FlagSet) {
	//load cfg
	configFile, err := cfg.LoadConfig(cfgFile, flagSet)
	if err != nil {
		panic(err)
	}

	//init logger
	log.New(cfg.AppConfig.AppDebug)
	defer log.Sync()

	log.Info("Using cfg file: " + configFile)
	log.Debug("cfg data", zap.String("config_data", helper.ToJsonString(cfg.AppConfig)))

	defer db.Close()
	db.InitConnections()

	redis.InitConnections()

	dingtalk.InitInstances()

	//运行 api server
	httpServer := web.New()
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Error("api server ListenAndServe error, "+err.Error(), zap.Error(err))
		}
	}()

	log.Info("http server run success, listen: " + cfg.AppConfig.HttpServer.Address)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	//todo 增加平滑重启信号相关处理逻辑
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
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
