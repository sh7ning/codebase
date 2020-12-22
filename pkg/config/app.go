package config

import "codebase/pkg/log"

type AppConfig struct {
	AppName  string            `mapstructure:"app_name"`
	AppDebug bool              `mapstructure:"app_debug"`
	Logger   *log.LoggerConfig `mapstructure:"logging" validate:"required,dive"`
}

func (appCfg *AppConfig) InitLoggerConfig() {
	appCfg.Logger.AppName = appCfg.AppName
	appCfg.Logger.Development = appCfg.AppDebug
}
