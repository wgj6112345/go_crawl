package main

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/proxy/logger"
	"github.com/wgj6112345/go_crawl/proxy/schedular"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

func main() {
	s := schedular.Schedular{}
	s.Init()

	http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		ip, err := getIP(s)
		if err != nil {
			logger.Logger.Errorf("getip failed, err : %v\n", err)
			return
		}
		fmt.Fprintf(w, ip)
	})

	http.HandleFunc("/collect", func(w http.ResponseWriter, r *http.Request) {
		
	})
	http.ListenAndServe(":9191", nil)
}

func getIP(s schedular.Schedular) (ip string, err error) {
	conn := s.RedisPool.Get()
	defer conn.Close()
	ip, err = redis.String(conn.Do("LPOP", s.ProxyQueueRedisKey))
	if err != nil {
		logger.Logger.Errorf("LPOP %s failed, err : %v\n", s.ProxyQueueRedisKey, err)
		return
	}

	return
}
