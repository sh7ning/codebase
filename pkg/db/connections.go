package db

import (
	"codebase/pkg/log"
	"fmt"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type Configs map[string]*Config

type Connections struct {
	configs     Configs
	collections map[string]*gorm.DB
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
