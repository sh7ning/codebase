package db

import (
	"app/pkg/cfg"
	"app/pkg/services/log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

var (
	dbList = sync.Map{}
)

func newDB(name string, config cfg.DB) (*gorm.DB, error) {
	db, err := gorm.Open(config.Type, config.DSN)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	db.LogMode(cfg.AppConfig.AppDebug)
	//db.SetLogger(log1.New(os.Stdout, "\r\n", 0))

	db.DB().SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`
	db.SingularTable(true)

	dbList.Store(name, db)
	//dbList[name] = db
	return db, nil
}

func Connection(conn string) *gorm.DB {
	if conn == "" {
		conn = "default"
	}

	db, ok := dbList.Load(conn)
	if !ok {
		if config, ok := cfg.AppConfig.DB[conn]; ok {
			var err error
			db, err = newDB(conn, config)
			if err != nil {
				log.Panic("不存在的 db: "+conn, zap.Error(err))
			}
		} else {
			log.Error("不存在的 db 配置: " + conn)

			return nil
		}
	}

	return db.(*gorm.DB)
}
