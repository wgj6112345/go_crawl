package main

import (
	"flag"
	"fmt"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/rpc/ElasticService/service"

	"github.com/wgj6112345/go_crawl/rpc/rpcHelper"
)

var (
	elastic_host = flag.String("es", "", "elastic_host")
)

func main() {
	flag.Parse()

	logger.Logger.Infof("ElasticService serve host: %v\n", *elastic_host)
	serve(fmt.Sprintf(":%s", *elastic_host))
}

func serve(host string) {
	rpcHelper.ServeRPC(host, new(service.DemoSaveService))
}
