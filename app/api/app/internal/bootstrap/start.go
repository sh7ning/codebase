package bootstrap

import (
	"codebase/app/api/app/internal/global"
	"codebase/app/api/app/internal/web"
	"codebase/pkg/defers"
	"codebase/pkg/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

func Start(cfgFile string, flagSet *pflag.FlagSet) {
	defer defers.Run()

	//初始化全局变量，资源等
	global.Init(cfgFile, flagSet)

	//业务代码开始
	web.New()

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
}
