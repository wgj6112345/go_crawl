package client

import (
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model/book"
	"imooc/分布式爬虫项目/demo1/rpc/rpcHelper"
)

func SaveItem(host string) (out chan book.BookItem, err error) {
	out = make(chan book.BookItem)
	client := rpcHelper.NewClient(host)

	go func() {
		var count int
		for {
			item := <-out
			logger.Logger.Infof("got %v item: %v \n", count, item)
			count++

			// rpc 通信
			var result string
			err := client.Call("DemoSaveService.Save", item, &result)
			if err != nil {
				logger.Logger.Errorf("client call DemoSaveService.Save failed, err : %v\n", err)
				return
			}

			logger.Logger.Infof("DemoSaveService.Save save success, item: %v\n", result)
		}
	}()
	return
}
