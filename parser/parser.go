package parser

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/model"
	"github.com/wgj6112345/go_crawl/model/book"
	bookpkg "github.com/wgj6112345/go_crawl/model/book"
	"regexp"
)

// 大多数网站都是 三级页面
// Level1 : 整体的类别，比如  城市名：北京，上海，武汉等， 书籍类别：小说，历史，科幻等
// Level2 : 具体类别的列表， 比如  北京的酒店，武汉的酒店等 ； 比如 小说类的书籍列表，科幻类的书籍列表等
// Level3 : 详情页，比如 XXX酒店 详情页； 比如 《斗破苍穹》、《三体》的详情页
const (
	urlBase      = "https://book.douban.com"
	urlPrefix    = "https://book.douban.com/tag/"
	reLevel1     = `<a href="/tag/([^"]+)">([^<]+)</a>`
	reLevel2     = `<a href="([^"]+)" title="([^"]+)"`
	reLevel2Flag = `<span class="thispage">1</span>`
	reLevel2Page = `<a href="([^"]+)">\d</a>[\d\D]*?`
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

type parseFunc func([]byte) model.ParseResult

type Level12Parser struct {
	ParseFunc parseFunc
	Name      string
}

func (p *Level12Parser) Parse(body []byte) (parseResult model.ParseResult) {
	return p.ParseFunc(body)
}

func (p *Level12Parser) Serialize() (name string, args []interface{}) {
	return p.Name, nil
}

type Level3Parser struct {
	Url  string
	Cate string
}

func (p *Level3Parser) Parse(body []byte) (parseResult model.ParseResult) {
	return ParseLevel3(body, p.Url, p.Cate)
}

func (p *Level3Parser) Serialize() (name string, args []interface{}) {
	return "ParseLevel3", []interface{}{p.Url, p.Cate}
}

type NilParser struct{}

func (p *NilParser) Parse(body []byte) model.ParseResult {
	return model.ParseResult{}
}

func (p *NilParser) Serialize() (name string, args []interface{}) {
	return "NilParser", nil
}

func ParseLevel1(body []byte) (result model.ParseResult) {
	re := regexp.MustCompile(reLevel1)
	match := re.FindAllSubmatch(body, -1)

	// url 正则项位于第几个
	var urlIndex = 1
	// var itemIndex = 2
	var request model.Request
	for _, m := range match {
		request.Url = urlPrefix + string(m[urlIndex])
		request.Parser = &Level12Parser{
			ParseFunc: ParseLevel2,
			Name:      "ParseLevel2",
		}
		result.Requests = append(result.Requests, request)
		// result.Items = append(result.Items, string(m[itemIndex]))
	}
	return result
}

// level2 里面有很多页面，需要翻页
// 如果超过 9 页，取前9页
// 如果只有一页，就取 1 页
// 如果不足 9 页，取最大页数
func ParseLevel2(body []byte) (result model.ParseResult) {
	// curpage=1 才获取页码  避免重复获取 url
	re := regexp.MustCompile(reLevel2Flag)
	flag := re.FindAllSubmatch(body, -1)

	if flag != nil {
		// 获取页码 仅取前几页
		re := regexp.MustCompile(reLevel2Page)
		pages := re.FindAllSubmatch(body, -1)

		var request model.Request
		for _, pageUrl := range pages {
			page := fmt.Sprintf("%s%s", urlBase, string(pageUrl[1]))
			request.Url = page
			request.Parser = &Level12Parser{
				ParseFunc: ParseLevel2,
				Name:      "ParseLevel2",
			}
			result.Requests = append(result.Requests, request)
		}
	}

	// 获取图书类别
	reCate := `<h1>.*?: ([^<]+)</h1>`
	re = regexp.MustCompile(reCate)
	cate := re.FindSubmatch(body)

	re = regexp.MustCompile(reLevel2)
	match := re.FindAllSubmatch(body, -1)
	var request model.Request
	var urlIndex = 1
	// var itemIndex = 2
	for _, m := range match {
		request.Url = string(m[urlIndex])
		request.Parser = &Level3Parser{
			Url:  string(m[urlIndex]),
			Cate: string(cate[1]),
		}
		result.Requests = append(result.Requests, request)
		// result.Items = append(result.Items, string(m[itemIndex]))
	}

	// logger.Logger.Infof("parseLevel2: result: \n", result)
	return
}

func ParseLevel3(body []byte, url string, cate string) (result model.ParseResult) {
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

	var book = book.BookDetail{
		Cate:        cate,
		Name:        mName,
		Author:      mAuthor,
		Publisher:   mPublisher,
		PublishTime: mPublishTime,
		Price:       mPrice,
		Score:       mScore,
		Intro:       mIntro,
	}

	// url type id
	// 获取 url 中的 id
	reUrlId := `([\d]+)`
	reId := regexp.MustCompile(reUrlId)
	id := reId.FindString(url)

	var item bookpkg.BookItem
	item.Url = url
	item.Id = id
	item.BookDetail = book
	result.Items = append(result.Items, item)
	// logger.Logger.Infof("parseLevel2: result: \n", result)
	return
}

func NilParseer([]byte) model.ParseResult {
	return model.ParseResult{}
}

func findString(body []byte, re *regexp.Regexp) string {
	result := re.FindSubmatch(body)

	if len(result) >= 2 {
		return string(result[1])
	} else {
		return ""
	}
}
