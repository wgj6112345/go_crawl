package engine

import "imooc/分布式爬虫项目/demo1/model"

type Engine interface {
	Run(seeds ...model.Request)
}
