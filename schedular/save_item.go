package schedular

import (
	"fmt"
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model/book"
)

const (
	index = "douban"
	typ   = "book"

	elastic_host = ":10001"
)

func SaveItem() chan book.BookItem {
	out := make(chan book.BookItem)

	go func() {
		var count int
		for {
			item := <-out
			logger.Logger.Infof("got %v item: %v \n", count, item)
			count++

			// TODO 插入 es
			// 试了一个多小时 es 无法插入数据 不知道是哪里的问题 后面再试试
			result, err := DemoSave(item)
			if err != nil {
				logger.Logger.Errorf("save item id: %d failed, err : %v\n", item.Id, err)
				return
			}
			logger.Logger.Infof("save item : %v success\n", result)
		}
	}()
	return out
}

// 暂时无法往 es 插入数据 先用 demosave
func save(item book.BookItem) (err error) {
	return
}

func DemoSave(item book.BookItem) (result string, err error) {
	// logger.Logger.Infof("saving item: %v\n", item)

	return fmt.Sprintf("%v", item), nil
}
