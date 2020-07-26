package client

import (
	"imooc/分布式爬虫项目/demo1/engine"
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/rpc/WorkerService/service"
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
