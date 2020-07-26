package main

import (
	"flag"
	"fmt"
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/rpc/WorkerService/service"
	"imooc/分布式爬虫项目/demo1/rpc/rpcHelper"
)

var (
	worker_host = flag.String("work", "", "worker_host")
)

func main() {
	flag.Parse()

	logger.Logger.Infof("WokerService serve host: %v\n", *worker_host)
	serve(fmt.Sprintf(":%s", *worker_host))
}

func serve(host string) {
	rpcHelper.ServeRPC(host, new(service.WorkerService))
}
