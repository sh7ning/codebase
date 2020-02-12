package cfg

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var AppConfig app

type app struct {
	AppDebug   bool                `mapstructure:"app_debug"`
	App        App                 `mapstructure:"app" validate:"required,dive"`
	HttpServer HttpServer          `mapstructure:"http_server" validate:"required"`
	DB         map[string]DB       `mapstructure:"db" validate:"required,dive"`
	Redis      map[string]Redis    `mapstructure:"redis" validate:"dive"`
	DingTalk   map[string]DingTalk `mapstructure:"dingtalk"`
}

type App struct {
	Name           string `mapstructure:"name" validate:"required"`
	LogFile        string `mapstructure:"log_file"`
	ErrorReporting bool   `mapstructure:"error_reporting"`
}

type HttpServer struct {
	Address string `mapstructure:"address" validate:"required"`
	Token   string `mapstructure:"token"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" validate:"required"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

type DingTalk struct {
	Token  string `mapstructure:"token" validate:"required"`
	Secret string `mapstructure:"secret"`
}

type DB struct {
	Type string `mapstructure:"type" validate:"required"`
	DSN  string `mapstructure:"dsn" validate:"required"`

	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" validate:"required"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" validate:"required"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" validate:"required"`
}

func LoadConfig(cfgFile string) (string, error) {
	v := viper.New()
	v.SetConfigFile(cfgFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		return "", err
	}

	if err := v.Unmarshal(&AppConfig, func(config *mapstructure.DecoderConfig) {
		config.TagName = "mapstructure"
		//config.TagName = "yaml"
		// do anything your like
	}); err != nil {
		return "", err
	}

	//校验配置
	validate := validator.New()
	if err := validate.Struct(AppConfig); err != nil {
		return "", err
	}

	return v.ConfigFileUsed(), nil
}
