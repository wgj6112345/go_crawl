package main

import (
	"imooc/分布式爬虫项目/demo1/engine"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/parser"
	"imooc/分布式爬虫项目/demo1/schedular"
)

func main() {
	url := "https://book.douban.com/tag/"

	s := schedular.DefaultSchedular{}
	e := engine.NewConCurrentEngine(&s, 1000)
	e.Run(model.Request{
		Url:       url,
		ParseFunc: parser.ParseLevel1,
	})
}
