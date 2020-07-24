package schedular

import (
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	redisdb "imooc/分布式爬虫项目/demo1/proxy/redis"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	redis_addr    = "127.0.0.1:6379"
	max_idle      = 20
	max_active    = 120
	idle_timeout  = 1
	proxyPoolKey  = "proxy_pool"
	proxyQueueKey = "proxy_queue"
)

type Schedular struct {
	ipChan  chan model.IP
	outChan chan model.IP

	FlushConn          redis.Conn
	RedisPool          *redis.Pool
	ProxyQueueRedisKey string

	collectChan chan bool
	syncChan    chan bool
}

func (s *Schedular) Init() {
	s.ipChan = make(chan model.IP, 10000)
	s.outChan = make(chan model.IP, 10000)
	s.collectChan = make(chan bool)
	s.RedisPool = redisdb.InitRedis(redis_addr, max_idle, max_active, idle_timeout)
	s.ProxyQueueRedisKey = "proxy_queue"
}

func (s *Schedular) Run() {
	// 初始化
	s.Init()
	// 开协程收集
	go func() {
		for {
			<-s.collectChan

			logger.Logger.Infof("start recollect new proxy...")
			Collect(s.ipChan)
		}
	}()
	// 验证 ip 合法性 和 可用性
	for i := 0; i < 50; i++ {
		go s.Verify()
	}

	s.CheckIpNum()
	// 往 redis 里面存 TODO   布隆过滤器 判断 是否 redis 已经存在 ip
	// ZADD key speed value
	// s.SyncTORedis()
}

func (s *Schedular) Verify() {
	// ticker := time.NewTicker(time.Second * 3)
	var count int
	conn := s.RedisPool.Get()
	defer conn.Close()
	for {
		select {
		case ip := <-s.ipChan:
			if count > 2 {
				conn.Flush()
				count = 0
			}

			if verify(ip) {
				conn.Send("LPUSH", s.ProxyQueueRedisKey, ip.Ip)
				count++
				logger.Logger.Debugf("count: %v\n", count)
			}
			// case <-ticker.C:
			// 	logger.Logger.Debugf("定时刷新 pipeline...\n")
			// 	conn.Flush()
			// default:
			// 	// logger.Logger.Debugf("wait new proxy come...\n")
			// 	time.Sleep(time.Second * 2)
		}
	}
}

func (s *Schedular) CheckIpNum() {

	ticker := time.NewTicker(time.Second * 10)
	conn := s.RedisPool.Get()
	defer conn.Close()

	for {
		select {
		case <-ticker.C:
			// logger.Logger.Infof("check proxy len\n")
			len, err := redis.Int(conn.Do("LLEN", s.ProxyQueueRedisKey))
			if err != nil {
				logger.Logger.Errorf("llen %s failed, err : %v\n", s.ProxyQueueRedisKey, err)
				return
			}

			if len < 1000 {
				// logger.Logger.Infof("proxy len : %v\n", len)
				s.collectChan <- true
			}
		}
	}
}

// func (s *Schedular) SyncTORedis() {
// 	<-s.syncChan

// 	logger.Logger.Infof("start sync to redis\n")
// 	conn := s.RedisPool.Get()
// 	defer conn.Close()

// 	// 响应时间控制在 1s 内，其实不用排序
// 	// ipMap, err := redis.Values(conn.Do("zrevrange", s.ProxyPoolRedisKey, 0, -1))
// 	// if err != nil {
// 	// 	logger.Logger.Errorf("zrange %s err: %v\n", s.ProxyPoolRedisKey, err)
// 	// 	return
// 	// }

// 	// for _, ip := range ipMap {
// 	// 	conn.Send("LPUSH", s.ProxyQueueRedisKey, string(ip.([]byte)))
// 	// }

// 	conn.Flush()
// }
