package service

import (
	"fmt"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/parser"
	"imooc/分布式爬虫项目/demo1/rpc/config"
	"imooc/分布式爬虫项目/demo1/rpc/rpcHelper"
	"testing"
)

func TestProcessor(t *testing.T) {
	t.Skip()
	mReq := model.Request{
		Url: "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
		Parser: &parser.Level12Parser{
			ParseFunc: parser.ParseLevel2,
			Name:      "ParseLevel2",
		},
	}

	w := WorkerService{}

	var result ParseResult
	req := SerializeRequest(mReq)
	err := w.Process(req, &result)
	if err != nil {
		t.Errorf("process failed, err : %v\n", err)
		return
	}

	fmt.Println("result: ", result)
}

func TestWokerService(t *testing.T) {
	mReq := model.Request{
		Url: "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C",
		Parser: &parser.Level12Parser{
			ParseFunc: parser.ParseLevel2,
			Name:      "ParseLevel2",
		},
	}
	req := SerializeRequest(mReq)
	var result ParseResult
	client := rpcHelper.NewClient(config.WORKER_SERVICE_HOST)

	err := client.Call("WorkerService.Process", req, &result)
	if err != nil {
		t.Errorf("process failed, err : %v\n", err)
		return
	}
	fmt.Println("result: ", result)
}
