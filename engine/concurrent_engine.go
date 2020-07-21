package engine

import (
	"imooc/分布式爬虫项目/demo1/fetcher"
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/schedular"
)

type ConCurrentEngine struct {
	Schedular schedular.Schedular
	workNum   int
}

func NewConCurrentEngine(s schedular.Schedular, workNum int) *ConCurrentEngine {
	engine := &ConCurrentEngine{
		Schedular: s,
		workNum:   workNum,
	}
	engine.Schedular.Init()
	return engine
}

func (e *ConCurrentEngine) Run(seeds ...model.Request) {

	out := make(chan model.ParseResult, 10000)
	for i := 0; i < e.workNum; i++ {
		HandleWorker(e.Schedular.GetWorkChan(), out)
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

		for _, item := range parseResult.Items {
			logger.Logger.Infof("got item : %v \n", item)
		}
	}

}

func HandleWorker(in chan model.Request, out chan model.ParseResult) {
	go func() {
		for {
			req := <-in
			result, err := work(req)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func work(req model.Request) (result model.ParseResult, err error) {
	logger.Logger.Infof("fetch url: %v \n", req.Url)
	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		logger.Logger.Errorf("fetch url: %v failed, err : %v\n", req.Url, err)
		return
	}
	result = req.ParseFunc(body)
	return
}
