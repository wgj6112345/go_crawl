package main

import (
	"flag"
	"fmt"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/rpc/WorkerService/service"
	"github.com/wgj6112345/go_crawl/rpc/rpcHelper"
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
