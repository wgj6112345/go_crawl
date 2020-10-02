package model

import (
	"github.com/wgj6112345/go_crawl/model/book"
)

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []book.BookItem
}
