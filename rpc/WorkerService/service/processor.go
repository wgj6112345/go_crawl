package service

import (
	"github.com/wgj6112345/go_crawl/engine"
	"github.com/wgj6112345/go_crawl/logger"
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
