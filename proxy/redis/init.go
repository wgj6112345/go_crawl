package redis

import (
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"time"

	"github.com/garyburd/redigo/redis"
)

func InitRedis(addr string, maxIdle int, maxActive int, idleTimeout int) (pool *redis.Pool) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", addr)
			if err != nil {
				logger.Logger.Errorf("redis connect [%s] failed, err: %s\n", addr, err)
				panic(err)
			}
			return
		},
	}
	return
}
