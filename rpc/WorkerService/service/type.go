package service

import (
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model"
	"imooc/分布式爬虫项目/demo1/model/book"
	"imooc/分布式爬虫项目/demo1/parser"
	"strings"
)

type Request struct {
	Url             string
	SerializeParser SerializeParser
}

type ParseResult struct {
	Requests []Request
	Items    []book.BookItem
}

type SerializeParser struct {
	Name string
	Args []interface{}
}

func DeserializeRequest(req Request) model.Request {
	// parseLevel3 需要 两个参数 一个 url  一个 cate 类型
	var url string
	var cate string

	for _, arg := range req.SerializeParser.Args {
		argStr := arg.(string)
		if strings.Contains(argStr, "http") {
			url = argStr
		} else {
			cate = argStr
		}
	}

	var mParser model.Parser
	switch req.SerializeParser.Name {
	case "ParseLevel1":
		mParser = &parser.Level12Parser{
			ParseFunc: parser.ParseLevel1,
			Name:      "ParseLevel1",
		}
	case "ParseLevel2":
		mParser = &parser.Level12Parser{
			ParseFunc: parser.ParseLevel2,
			Name:      "ParseLevel2",
		}
	case "ParseLevel3":
		mParser = &parser.Level3Parser{
			Url:  url,
			Cate: cate,
		}
	case "NilParseer":
		mParser = &parser.NilParser{}
	default:
		logger.Logger.Errorf("unknow parse name!!!\n")
	}

	result := model.Request{
		Url:    req.Url,
		Parser: mParser,
	}

	return result
}

func SerializeRequest(mReq model.Request) Request {
	name, args := mReq.Parser.Serialize()
	return Request{
		Url: mReq.Url,
		SerializeParser: SerializeParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(mResult model.ParseResult) ParseResult {
	var result ParseResult

	for _, mReq := range mResult.Requests {
		result.Requests = append(result.Requests, SerializeRequest(mReq))
	}
	result.Items = mResult.Items

	return result
}

func DeserializeResult(result ParseResult) model.ParseResult {
	var mResult model.ParseResult

	for _, mReq := range result.Requests {
		mResult.Requests = append(mResult.Requests, DeserializeRequest(mReq))
	}

	mResult.Items = result.Items
	return mResult
}
