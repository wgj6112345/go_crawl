package db

import (
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model"
)

const (
	index = "douban"
	typ   = "book"
)

func SaveItem() chan model.BookItem {
	out := make(chan model.BookItem)

	go func() {
		var count int
		for {
			item := <-out
			logger.Logger.Infof("got %v item: %v \n", count, item)
			count++

			// TODO 插入 es
			// 试了一个多小时 es 无法插入数据 不知道是哪里的问题 后面再试试
			// err := save(item.Id, item.Payload)
			// if err != nil {
			// 	logger.Logger.Errorf("save item id: %d failed, err : %v\n", item.Id, item.Payload)
			// 	return
			// }
			// logger.Logger.Infof("save item id: %d success\n", item.Id)
		}
	}()
	return out
}

func save(id string, data interface{}) (err error) {
	return Insert(index, typ, id, data)
}
