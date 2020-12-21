package config

import "codebase/pkg/log"

type AppConfig struct {
	AppName      string            `mapstructure:"app_name"`
	AppDebug     bool              `mapstructure:"app_debug"`
	LoggerConfig *log.LoggerConfig `mapstructure:"logging" validate:"required,dive"`
}

func (appCfg *AppConfig) InitLoggerConfig() {
	appCfg.LoggerConfig.AppName = appCfg.AppName
	appCfg.LoggerConfig.Development = appCfg.AppDebug
}
