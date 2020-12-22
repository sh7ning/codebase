package config

type AppConfig struct {
	AppName  string `mapstructure:"app_name"`
	AppDebug bool   `mapstructure:"app_debug"`
	Logger   struct {
		File string `mapstructure:"file"`
	} `mapstructure:"logging" validate:"required,dive"`
}
