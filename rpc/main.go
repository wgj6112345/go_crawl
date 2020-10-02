package main

import (
	"flag"
	"fmt"
	"github.com/wgj6112345/go_crawl/engine"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/model"
	"github.com/wgj6112345/go_crawl/parser"
	esClient "github.com/wgj6112345/go_crawl/rpc/ElasticService/client"
	workerClient "github.com/wgj6112345/go_crawl/rpc/WorkerService/client"
	"github.com/wgj6112345/go_crawl/rpc/rpcHelper"
	"github.com/wgj6112345/go_crawl/schedular"
	"net/rpc"
	"strings"
)

var (
	elastic_host = flag.String("es", "", "elastic_host")
	worker_hosts = flag.String("works", "", "workers_host")
)

func main() {
	flag.Parse()

	url := "https://book.douban.com/tag/"
	// url := "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C"
	// url := "https://book.douban.com/subject/30293801/"

	itemChan, err := esClient.SaveItem(fmt.Sprintf(":%s", *elastic_host))
	if err != nil {
		logger.Logger.Errorf("client.SaveItem failed, err: %v\n", err)
		return
	}

	var hosts []string
	worker_hosts := strings.Split(*worker_hosts, ",")

	for _, worker_host := range worker_hosts {
		hosts = append(hosts, fmt.Sprintf(":%s", worker_host))
	}

	clientPool := createClientPool(hosts)
	processor := workerClient.Processor(clientPool)

	s := schedular.QueueSchedular{}
	e := engine.ConCurrentEngine{
		Schedular:   &s,
		WorkNum:     100,
		ItemChan:    itemChan,
		ProcessFunc: processor,
	}

	e.Run(model.Request{
		Url: url,
		Parser: &parser.Level12Parser{
			ParseFunc: parser.ParseLevel1,
			Name:      "ParseLevel1",
		},
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clientPool []*rpc.Client

	for _, host := range hosts {
		client := rpcHelper.NewClient(host)
		logger.Logger.Infof("connect host: %v success\n", host)
		clientPool = append(clientPool, client)
	}

	out := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clientPool {
				out <- client
			}
		}
	}()

	return out
}
