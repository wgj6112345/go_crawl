package main

import (
	"imooc/分布式爬虫项目/demo1/db"
	"imooc/分布式爬虫项目/demo1/engine"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/parser"
	"imooc/分布式爬虫项目/demo1/schedular"
)

func main() {
	url1 := "https://book.douban.com/tag/"
	// url2 := "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C"
	// url3 := "https://book.douban.com/subject/30293801/"

	s := schedular.QueueSchedular{}
	e := engine.NewConCurrentEngine(&s, 100)
	e.ItemChan = db.SaveItem()

	e.Run(model.Request{
		Url:       url1,
		ParseFunc: parser.ParseLevel1,
	})
}
