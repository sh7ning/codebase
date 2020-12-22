package signal

import (
	"codebase/pkg/log"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func Wait() {
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	//todo 增加平滑重启信号相关处理逻辑
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	sig := <-quit
	log.Warn("get signal, start shutdown server ...", zap.Any("signal", sig))
}
