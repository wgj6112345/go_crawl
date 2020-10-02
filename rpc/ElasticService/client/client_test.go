package client

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/logger"
	"github.com/wgj6112345/go_crawl/model/book"
	"github.com/wgj6112345/go_crawl/rpc/rpcHelper"
	"testing"
)

const host = ":10001"

func TestDemoSaveService(t *testing.T) {
	var bookItem book.BookItem
	var result string

	bookItem = book.BookItem{
		Url:  "www.baidu.com",
		Type: "type",
		Id:   "3021",
		BookDetail: book.BookDetail{
			Cate:        "文学",
			Name:        "《局外人》",
			Author:      "王高杰",
			Publisher:   "上海译文出版社",
			PublishTime: "2010-8",
			Price:       "22.00 元",
			Score:       "9.0",
			Intro:       "《局外人》是加缪小说的成名作和代表作之一，堪称20世纪整个西方文坛最具有划时代意义最著名小说之一，“局外人”也由此成为整个西方文学-哲学中最经典的人物形象和最重要的关键词之一。",
		},
	}

	client := rpcHelper.NewClient(host)
	err := client.Call("DemoSaveService.Save", bookItem, &result)
	if err != nil {
		logger.Logger.Errorf("client call DemoSaveService.Save failed, err : %v\n", err)
		return
	}

	fmt.Println("result: ", result)
}
