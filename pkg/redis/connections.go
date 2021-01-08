package redis

import (
	"codebase/pkg/log"
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type Configs map[string]*Config

type Connections struct {
	configs     Configs
	collections map[string]*redis.Client
}

var connections *Connections

func Init(configs Configs) (err error) {
	connections, err = New(configs)
	return
}

func Conn(conn string) *redis.Client {
	return connections.Get(conn)
}

func New(configs Configs) (*Connections, error) {
	collections := make(map[string]*redis.Client)
	for conn, config := range configs {
		redisPool, err := NewRedis(config)
		if err != nil {
			return nil, fmt.Errorf("NewRedis error, conn: %s, error: %s", conn, err.Error())
		}

		collections[conn] = redisPool
	}

	return &Connections{
		configs:     configs,
		collections: collections,
	}, nil
}

func (c *Connections) Get(conn string) *redis.Client {
	if conn == "" {
		conn = "default"
	}

	if cli, ok := c.collections[conn]; ok {
		return cli
	}

	return nil
}

func (c *Connections) Publish(conn string, channel string, message interface{}) error {
	if err := c.Get(conn).Publish(channel, message).Err(); err != nil {
		log.Error("redis publish error, channel: "+channel, zap.Error(err), zap.Any("message", message))
		return err
	}

	log.Debug("redis publish channel:"+channel, zap.Any("message", message))
	return nil
}
