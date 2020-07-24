package schedular

import "imooc/分布式爬虫项目/demo1/model"

// 调度器 根据调度规则不同 可以另外实现 负载均衡等
type Schedular interface {
	Dispatch(req model.Request)
	GetWorkChan() chan model.Request
	WorkerIdle(chan model.Request)
	Run()
}
