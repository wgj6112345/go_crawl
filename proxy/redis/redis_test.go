package redis

import (
	"fmt"
	"testing"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	redis_addr   = "127.0.0.1:6379"
	max_idle     = 20
	max_active   = 80
	idle_timeout = 1
)

func TestFlush(t *testing.T) {
	pool := InitRedis(redis_addr, max_idle, max_active, idle_timeout)
	go func() {
		conn := pool.Get()
		for i := 0; i < 100; i++ {
			conn.Send("LPUSH", "rand", i)
		}
		conn.Flush()
	}()

	time.Sleep(time.Second * 10)

	conn := pool.Get()
	conn.Flush()
}

func TestZadd(t *testing.T) {
	t.Skip()
	pool := InitRedis(redis_addr, max_idle, max_active, idle_timeout)
	conn := pool.Get()

	type user struct {
		age  int
		name string
	}

	users := []user{{27, "wanggaojie"}, {25, "qijiamiao"}, {29, "feige"}}
	// age 27 wanggaojie 25 qijiamiao 29 feige
	// var item string
	// for index, user := range users {
	// 	if index == 1 {
	// 		item = fmt.Sprintf("%d %s", user.age, user.name)
	// 		continue
	// 	}
	// 	item += fmt.Sprintf(",%d %s", user.age, user.name)
	// }

	// fmt.Println("item: ", item)

	for _, user := range users {
		conn.Send("zadd", "proxy_test", user.age, user.name)
	}

	conn.Flush()

	userMap, err := redis.Values(conn.Do("zrevrange", "proxy_test", 0, -1))
	if err != nil {
		fmt.Println("zrange err: ", err)
		return
	}

	for _, user := range userMap {
		fmt.Println("user: ", user.([]byte))
		conn.Send("LPUSH", "proxy_queue", string(user.([]byte)))
	}

	conn.Flush()
}
