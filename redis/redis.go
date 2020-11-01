package redis

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	addr        string = "127.0.0.1:6379"
	maxIdle     int    = 10
	maxActive   int    = 20
	idleTimeout int    = 180
)

func InitRedis() (pool *redis.Pool) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", addr)
			if err != nil {
				fmt.Printf("redis connect [%s] failed, err: %s\n", addr, err)
				panic(err)
			}
			return
		},
	}
	return
}
