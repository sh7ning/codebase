package app

import (
	"codebase/pkg/config"
	"codebase/pkg/helper"
	"codebase/pkg/log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/pflag"
	"go.uber.org/zap"
)

type ConfigInterface interface {
	AppConfig() *Config
	mustEmbed()
}

type Config struct {
	AppName  string `mapstructure:"app_name"`
	AppDebug bool   `mapstructure:"app_debug"`
	Logger   struct {
		File string `mapstructure:"file"`
	} `mapstructure:"logging" validate:"required,dive"`
}

func (c *Config) AppConfig() *Config {
	return c
	//return nil, errors.New("method AppConfig not implemented")
}

func (*Config) mustEmbed() {}

type App struct {
	errs chan error
}

func New(cfg ConfigInterface, cfgFile string, flagSet *pflag.FlagSet) (*App, error) {
	//load cfg
	c, err := config.LoadConfig(cfg, cfgFile, flagSet, nil)
	if err != nil {
		return nil, err
	}

	//init logger
	log.New(&log.LoggerConfig{
		Development: cfg.AppConfig().AppDebug,
		AppName:     cfg.AppConfig().AppName,
		LogFile:     cfg.AppConfig().Logger.File,
		Notify:      nil,
	})

	log.Info("using cfg file: " + c.ConfigFileUsed())
	log.Debug("cfg data", zap.String("config_data", helper.ToJsonString(cfg)))
	return &App{
		errs: make(chan error, 10),
	}, nil
}

func (a *App) RunWith(boot func(chan error) error) error {
	if err := boot(a.errs); err != nil {
		return err
	}

	return a.wait()
}

func (a *App) wait() error {
	// Wait for interrupt signal to gracefully shutdown the server with
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	//todo 增加平滑重启信号相关处理逻辑
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	select {
	case err := <-a.errs:
		return err

	case sig := <-quit:
		log.Warn("get signal, start shutdown server ...", zap.Any("signal", sig))
		return nil
	}
}
