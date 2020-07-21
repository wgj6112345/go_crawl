package engine

import (
	"imooc/分布式爬虫项目/demo1/logger"
	"imooc/分布式爬虫项目/demo1/model"
	"regexp"
)

const (
	urlPrefix = "https://book.douban.com/tag/"
	reLevel1  = `<a href="/tag/([^"]+)">([^<]+)</a>`
	reLevel2  = `<a href="([^"]+)" title="([^"]+)"`
)

// reLevel3
const (
	reLevel3    = ""
	name        = `<h1>[\d\D]*?<span.*?>([^<]+)</span>[\d\D]*?</h1>`
	author      = `<span.*?> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`
	publisher   = `<span.*?>出版社:</span> ([^<]+)<`
	publishTime = `<span.*?>出版年:</span> ([^<]+)<`
	price       = `<span.*?>定价:</span> ([^<]+)<`
	score       = `<strong class="ll rating_num ".*?> ([^<]+) </strong>`
	intro       = `<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`
)

func ParseLevel1(body []byte) (result ParseResult) {
	re := regexp.MustCompile(reLevel1)
	match := re.FindAllSubmatch(body, -1)

	// url 正则项位于第几个
	var urlIndex = 1
	var itemIndex = 2
	var request Request
	for _, m := range match {
		request.Url = urlPrefix + string(m[urlIndex])
		request.ParseFunc = ParseLevel2
		result.Requests = append(result.Requests, request)
		result.Items = append(result.Items, string(m[itemIndex]))
	}
	return result
}

func ParseLevel2(body []byte) (result ParseResult) {
	re := regexp.MustCompile(reLevel2)
	match := re.FindAllSubmatch(body, -1)

	var urlIndex = 1
	var itemIndex = 2
	var request Request
	for _, m := range match {
		request.Url = string(m[urlIndex])
		request.ParseFunc = ParseLevel3
		result.Requests = append(result.Requests, request)
		result.Items = append(result.Items, string(m[itemIndex]))
	}
	// logger.Logger.Infof("parseLevel2: result: \n", result)
	return
}

func ParseLevel3(body []byte) (result ParseResult) {
	// 预编译
	reName := regexp.MustCompile(name)
	reAuthor := regexp.MustCompile(author)
	rePublisher := regexp.MustCompile(publisher)
	rePublishTime := regexp.MustCompile(publishTime)
	rePrice := regexp.MustCompile(price)
	reScore := regexp.MustCompile(score)
	reIntro := regexp.MustCompile(intro)

	mName := findString(body, reName)
	mAuthor := findString(body, reAuthor)
	mPublisher := findString(body, rePublisher)
	mPublishTime := findString(body, rePublishTime)
	mPrice := findString(body, rePrice)
	mScore := findString(body, reScore)
	mIntro := findString(body, reIntro)

	var book = model.BookDetail{
		Name:        mName,
		Author:      mAuthor,
		Publisher:   mPublisher,
		PublishTime: mPublishTime,
		Price:       mPrice,
		Score:       mScore,
		Intro:       mIntro,
	}

	result.Items = append(result.Items, book)
	// logger.Logger.Infof("parseLevel2: result: \n", result)
	return
}

func ParseLevel4(body []byte) (result ParseResult) {
	return
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

func ParserWithRe(body []byte, re *regexp.Regexp) {
	match := re.FindAllSubmatch(body, -1)
	logger.Logger.Infof("test parse: result: %v\n", match)
	return
}

func findString(body []byte, re *regexp.Regexp) string {
	result := re.FindSubmatch(body)

	if len(result) >= 2 {
		return string(result[1])
	} else {
		return ""
	}
}
