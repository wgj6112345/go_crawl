package main

import (
	"imooc/分布式爬虫项目/demo1/engine"
)

func main() {
	url := "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C"
	engine.Run(engine.Request{
		Url:       url,
		ParseFunc: engine.ParseLevel2,
	})
}
