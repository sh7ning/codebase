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

func Init(configs Configs) *Connections {
	collections := make(map[string]*redis.Client)
	for conn, config := range configs {
		redisPool, err := NewRedis(config)
		if err != nil {
			log.Panic(fmt.Sprintf("newRedis redis, conn: %s, error: %s", conn, err.Error()), zap.Error(err))
		}

		collections[conn] = redisPool
	}

	return &Connections{
		configs:     configs,
		collections: collections,
	}
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
