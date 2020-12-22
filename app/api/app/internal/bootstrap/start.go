package bootstrap

import (
	"codebase/app/api/app/internal/global"
	"codebase/app/api/app/internal/web"
	"codebase/pkg/defers"
	"codebase/pkg/signal"

	"github.com/spf13/pflag"
)

func Start(cfgFile string, flagSet *pflag.FlagSet) {
	defer defers.Run()

	//初始化全局变量，资源等
	global.Init(cfgFile, flagSet)

	//业务代码开始
	web.New()

	// Wait for interrupt signal to gracefully shutdown the server with
	signal.Wait()
}
