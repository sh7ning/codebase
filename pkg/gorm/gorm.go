package gorm

import (
	"codebase/pkg/log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

var dbLogger = &logger{}

type logger struct {
}

func (l logger) Print(v ...interface{}) {
	log.Debug("db log", zap.Any("v", v))
}

type Config struct {
	Type string `mapstructure:"type" validate:"required"`
	DSN  string `mapstructure:"dsn" validate:"required"`

	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" validate:"required"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" validate:"required"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" validate:"required"`
}

func NewDB(debug bool, config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(config.Type, config.DSN)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	db.LogMode(debug)
	db.SetLogger(dbLogger)

	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	db.SingularTable(true)

	return db, nil
}
