package redis

import (
	"app/pkg/cfg"
	"app/pkg/services/log"
	"errors"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var redisConnections = make(map[string]*redis.Client)

func newRedis(conn string, addr string, password string, db int) (*redis.Client, error) {
	log.Info("连接 redis: " + addr)
	redisPool := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
		//DialTimeout:  10 * time.Second,
		//ReadTimeout:  30 * time.Second,
		//WriteTimeout: 30 * time.Second,
		//PoolSize:     10,
		//PoolTimeout:  30 * time.Second,
	})

	if err := redisPool.Ping().Err(); err != nil {
		return nil, errors.New("redis 连接失败: " + err.Error())
	}

	return redisPool, nil
}

func InitConnections() {
	for conn, config := range cfg.AppConfig.Redis {
		redisPool, err := newRedis(conn, config.Addr, config.Password, config.Db)
		if err != nil {
			log.Panic("不存在的 redis: " + conn)
		}

		redisConnections[conn] = redisPool
	}
}

func Connection(conn string) *redis.Client {
	if conn == "" {
		conn = "default"
	}

	return redisConnections[conn]
}

func Publish(conn string, channel string, message interface{}) error {
	if err := Connection(conn).Publish(channel, message).Err(); err != nil {
		log.Error("redis publish error, channel: "+channel, zap.Error(err), zap.Any("message", message))
		return err
	}

	log.Debug("redis publish channel:"+channel, zap.Any("message", message))
	return nil
}
