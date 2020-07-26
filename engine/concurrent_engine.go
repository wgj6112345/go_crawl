package engine

import (
	"imooc/分布式爬虫项目/demo1/fetcher"
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/model/book"
	"imooc/分布式爬虫项目/demo1/schedular"
)

type Processor func(model.Request) (model.ParseResult, error)

type ConCurrentEngine struct {
	Schedular   schedular.Schedular
	WorkNum     int
	ItemChan    chan book.BookItem
	ProcessFunc Processor
}

func (e *ConCurrentEngine) Run(seeds ...model.Request) {
	e.Schedular.Run()

	out := make(chan model.ParseResult)
	for i := 0; i < e.WorkNum; i++ {
		e.HandleWorker(e.Schedular.GetWorkChan(), out, e.Schedular)
	}

	// 任务过来了 进行调度
	for _, req := range seeds {
		e.Schedular.Dispatch(req)
	}

	// 打印结果
	for {
		parseResult := <-out

		for _, req := range parseResult.Requests {
			e.Schedular.Dispatch(req)
		}

		// TODO 用于存储服务
		for _, item := range parseResult.Items {
			go func(item book.BookItem) { e.ItemChan <- item }(item)
		}
	}

}

// 改造 队列式的 worker 每个都有一个 work channel
func (e *ConCurrentEngine) HandleWorker(in chan model.Request, out chan model.ParseResult, s schedular.Schedular) {
	go func() {
		for {
			s.WorkerIdle(in)
			req := <-in
			result, err := e.ProcessFunc(req)
			if err != nil {
				continue
			}
			// TODO  插入 redis 布隆过滤器

			out <- result
		}
	}()
}

func Work(req model.Request) (result model.ParseResult, err error) {
	// logger.Logger.Infof("fetch url: %v \n", req.Url)

	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		logger.Logger.Errorf("fetch url: %v failed, err : %v\n", req.Url, err)
		return Work(req)
	}
	result = req.Parser.Parse(body)
	return
}
