package service

import (
	"imooc/分布式爬虫项目/demo1/engine"
	"imooc/分布式爬虫项目/demo1/logger"
)

type WorkerService struct {
}

func (s *WorkerService) Process(req Request, parseResult *ParseResult) (err error) {
	logger.Logger.Infof("process url: %v\n", req.Url)

	mReq := DeserializeRequest(req)
	mResult, err := engine.Work(mReq)
	if err != nil {
		logger.Logger.Errorf("engine.Worke request: %v failed, err: %v\n", req, err)
		return
	}

	*parseResult = SerializeResult(mResult)
	return
}
