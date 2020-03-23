package db

import (
	"app/pkg/cfg"
	"app/pkg/services/log"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

var (
	dbList   = make(map[string]*gorm.DB)
	dbLogger = &logger{}
)

type logger struct {
}

func (l logger) Print(v ...interface{}) {
	log.Debug("db log", zap.Any("v", v))
}

func newDB(name string, config cfg.DB) (*gorm.DB, error) {
	db, err := gorm.Open(config.Type, config.DSN)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	db.LogMode(cfg.AppConfig.AppDebug)
	db.SetLogger(dbLogger)

	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	db.SingularTable(true)

	return db, nil
}

func InitConnections() {
	for conn, config := range cfg.AppConfig.DB {
		db, err := newDB(conn, config)
		if err != nil {
			log.Panic("不存在的 db: "+conn, zap.Error(err))
		}
		dbList[conn] = db
	}
}

func Close() {
	for conn, db := range dbList {
		if err := db.Close(); err != nil {
			log.Error(fmt.Sprintf("db: %s close error: %s", conn, err.Error()))
		}
	}
}

func Connection(conn string) *gorm.DB {
	if conn == "" {
		conn = "default"
	}

	return dbList[conn]
}
