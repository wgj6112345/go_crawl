package model

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []BookItem
}

type BookItem struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
