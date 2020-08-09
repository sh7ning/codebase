package gorm

import (
	"codebase/pkg/log"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Configs map[string]*Config

type Connections struct {
	configs     Configs
	connections map[string]*gorm.DB
}

func InitConnections(debug bool, configs Configs) *Connections {
	connections := make(map[string]*gorm.DB)
	for name, config := range configs {
		db, err := NewDB(debug, config)
		if err != nil {
			log.Panic(fmt.Sprintf("NewDB error, name: %s, error: %s", name, err.Error()), zap.Error(err))
		}
		connections[name] = db
	}

	return &Connections{
		configs:     configs,
		connections: connections,
	}
}

func (conns *Connections) Close() {
	for name, db := range conns.connections {
		if err := db.Close(); err != nil {
			log.Error(fmt.Sprintf("db: %s close error: %s", name, err.Error()))
		}
	}
}

func (conns *Connections) Connection(conn string) *gorm.DB {
	if conn == "" {
		conn = "default"
	}

	return conns.connections[conn]
}
