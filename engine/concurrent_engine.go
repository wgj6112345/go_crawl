package engine

import (
	"fmt"
	"strings"

	"github.com/wgj6112345/go_crawl/fetcher"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/model"
	"github.com/wgj6112345/go_crawl/model/book"
	"github.com/wgj6112345/go_crawl/schedular"
	"github.com/wgj6112345/go_crawl/selenium"
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
			// logger.Logger.Infof("start save item: %v \n", item)
			go func(item book.BookItem) { e.ItemChan <- item }(item)
		}
	}

}

// 队列式 worker 每个都有一个 work channel
func (e *ConCurrentEngine) HandleWorker(in chan model.Request, out chan model.ParseResult, s schedular.Schedular) {
	go func() {
		for {
			s.WorkerIdle(in)
			req := <-in
			result, err := e.ProcessFunc(req)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

func Work(req model.Request) (result model.ParseResult, err error) {
	logger.Logger.Infof("fetch url: %v \n", req.Url)

	// 如果 url 在重复，则不作处理
	// 需要 redis 4.0以上，并且安装 bloomfilter 插件
	// redisPool := redis.InitRedis()
	// redisClient := redisPool.Get()

	// isExist, err := redis.Int(redisClient.Do("bf.exists", "url", req.Url))
	// if err != nil {
	// 	fmt.Println("bf.exists err: ", err)
	// 	return model.ParseResult{}, err
	// }

	// if isExist == 0 {
	// 	if _, err := redisClient.Do("bf.add", "url", req.Url); err != nil {
	// 		fmt.Println("bf.add err: ", err)
	// 		return model.ParseResult{}, err
	// 	}
	// } else if isExist == 1 {
	// 	return model.ParseResult{}, errors.New("重复url!!")
	// }

	body, err := fetcher.Fetch(req.Url)
	if err != nil {
		logger.Logger.Errorf("fetch url: %v failed, err : %v\n", req.Url, err)
		return Work(req)
	}

	// fmt.Println("body: ", body)

	result = req.Parser.Parse(body)
	// logger.Logger.Infof("result: %v \n", string(body))
	return
}

func WorkBySelenium(req model.Request) (result model.ParseResult, err error) {
	var (
		crawler *selenium.Crawler
		body    []byte
	)

	// logger.Logger.Infof("fetch url: %v \n", req.Url)

	if crawler, err = selenium.NewCrawler(); err != nil {
		logger.Logger.Errorf("selenium.NewCrawler failed, err : %v\n", err)
		return
	}

	if body, err = crawler.Fetch(req.Url); err != nil {
		logger.Logger.Errorf("fetch url: %v failed, err : %v\n", req.Url, err)
		return WorkBySelenium(req)
	}

	// fmt.Println("body: ", string(body))
	if strings.Contains(string(body), "网络不给力") {
		err = fmt.Errorf("被封了\n")
		return
	}

	result = req.Parser.Parse(body)
	logger.Logger.Infof("result: %v \n", result)
	return
}
