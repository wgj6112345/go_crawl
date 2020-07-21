package engine

import (
	"imooc/分布式爬虫项目/demo1/fetcher"
	"imooc/分布式爬虫项目/demo1/logger"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]

		logger.Logger.Infof("fetch url: %v \n", req.Url)
		body, err := fetcher.Fetch(req.Url)
		if err != nil {
			logger.Logger.Errorf("fetch url: %v failed, err : %v\n", req.Url, err)
			return
		}

		// logger.Logger.Debugf("body: %v\n", body)
		parseResult := req.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			logger.Logger.Infof("got item : %v \n", item)
		}

	}
}
