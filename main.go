package main

import (
	"imooc/分布式爬虫项目/demo1/engine"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/parser"
	"imooc/分布式爬虫项目/demo1/schedular"
)

func main() {
	// url := "https://book.douban.com/tag/"
	url := "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C"
	// url := "https://book.douban.com/subject/30293801/"

	s := schedular.QueueSchedular{}
	e := engine.ConCurrentEngine{
		Schedular:   &s,
		WorkNum:     1,
		ItemChan:    schedular.SaveItem(),
		ProcessFunc: engine.Work,
	}

	e.Run(model.Request{
		Url: url,
		Parser: &parser.Level12Parser{
			ParseFunc: parser.ParseLevel2,
			Name:      "ParseLevel2",
		},
	})
}
