package db

import (
	"codebase/pkg/log"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Configs map[string]*Config

type Connections struct {
	configs     Configs
	collections map[string]*gorm.DB
}

var dbs *Connections

func Init(debug bool, configs Configs) {
	dbs = New(debug, configs)
}

func Conn(conn string) *gorm.DB {
	return dbs.Get(conn)
}

func New(debug bool, configs Configs) *Connections {
	collections := make(map[string]*gorm.DB)
	for name, config := range configs {
		db, err := NewDB(debug, config)
		if err != nil {
			log.Panic(fmt.Sprintf("NewDB error, name: %s, error: %s", name, err.Error()), zap.Error(err))
		}
		collections[name] = db
	}

	return &Connections{
		configs:     configs,
		collections: collections,
	}
}

func (c *Connections) Get(conn string) *gorm.DB {
	if conn == "" {
		conn = "default"
	}

	if obj, ok := c.collections[conn]; ok {
		return obj
	}

	return nil
}
