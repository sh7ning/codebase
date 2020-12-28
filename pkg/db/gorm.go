package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Type string `mapstructure:"type" validate:"required"`
	DSN  string `mapstructure:"dsn" validate:"required"`

	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime" validate:"required"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns" validate:"required"`
	MaxOpenConns    int           `mapstructure:"max_open_conns" validate:"required"`
}

func NewDB(debug bool, config *Config) (*gorm.DB, error) {
	dial := mysql.Open(config.DSN)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: &logger{},
		NamingStrategy: schema.NamingStrategy{
			// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if config.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(config.ConnMaxIdleTime)
	}
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)

	return db, nil
}
