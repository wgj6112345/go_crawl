package main

import (
	"flag"
	"fmt"
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/rpc/ElasticService/service"

	"imooc/分布式爬虫项目/demo1/rpc/rpcHelper"
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
