package model

import (
	"imooc/分布式爬虫项目/demo1/model/book"
)

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []book.BookItem
}
