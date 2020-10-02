package client

import (
	"github.com/wgj6112345/go_crawl/engine"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/model"
	"github.com/wgj6112345/go_crawl/rpc/WorkerService/service"
	"net/rpc"
)

func Processor(clientPool chan *rpc.Client) engine.Processor {

	return func(mReq model.Request) (mResult model.ParseResult, err error) {
		req := service.SerializeRequest(mReq)

		var result service.ParseResult
		client := <-clientPool
		err = client.Call("WorkerService.Process", req, &result)
		if err != nil {
			logger.Logger.Errorf("process failed, err : %v\n", err)
			return
		}

		mResult = service.DeserializeResult(result)
		return
	}
}
