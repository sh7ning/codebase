package db

import (
	"fmt"

	"gorm.io/gorm"
)

type Configs map[string]*Config

type Connections struct {
	configs     Configs
	collections map[string]*gorm.DB
}

var dbs *Connections

func Init(debug bool, configs Configs) (err error) {
	dbs, err = New(debug, configs)
	return
}

func Conn(conn string) *gorm.DB {
	return dbs.Get(conn)
}

func New(debug bool, configs Configs) (*Connections, error) {
	collections := make(map[string]*gorm.DB)
	for name, config := range configs {
		db, err := NewDB(debug, config)
		if err != nil {
			return nil, fmt.Errorf("NewDB error, name: %s, error: %s", name, err.Error())
		}
		collections[name] = db
	}

	return &Connections{
		configs:     configs,
		collections: collections,
	}, nil
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
