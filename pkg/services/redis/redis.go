package redis

import (
	"app/pkg/cfg"
	"app/pkg/services/log"
	"errors"

	"github.com/go-redis/redis"
)

var redisConnections = make(map[string]*redis.Client)

func newRedis(conn string, addr string, password string, db int) error {
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
		return errors.New("redis 连接失败: " + err.Error())
	}

	redisConnections[conn] = redisPool
	return nil
}

func Connection(conn string) *redis.Client {
	if conn == "" {
		conn = "default"
	}

	if _, ok := redisConnections[conn]; !ok {
		if config, ok := cfg.AppConfig.Redis[conn]; ok {
			if err := newRedis(conn, config.Addr, config.Password, config.Db); err != nil {
				log.Panic("不存在的 redis: " + conn)
			}
		} else {
			log.Error("不存在的 redis 配置: " + conn)

			return nil
		}
	}

	return redisConnections[conn]
}
