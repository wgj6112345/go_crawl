package baidu

import (
	"github.com/wgj6112345/go_crawl/model"
	baidumodel "github.com/wgj6112345/go_crawl/model/baidu"
	"regexp"
	"strings"
)

var (
	re_baidu_selenium = `<div[\d\D]*?class="result c-container "[\d\D]*?href="([^"]*?)"[\d\D]*?>([\d\D]*?)</a></h3>[\d\D]*?<div class="c-abstract"><span class=" newTimeFactor_before_abs  m">([^<]*?)&nbsp;-&nbsp;</span>([\d\D]*?)</div>`

	re_baidu = `<div class="result c-container "[\d\D]*?href = "([^"]*?)"[\d\D]*?>([\d\D]*?)</a></h3>[\d\D]*?<div class="c-abstract"><span class=" newTimeFactor_before_abs  m">([^<]*?)&nbsp;-&nbsp;</span>([\d\D]*?)</div>`
	// re_baidu = `<div class="result c-container "[\d\D]*?href="([^"]+)">(.*?)</a></h3><div class="c-abstract"><span class=" newTimeFactor_before_abs  m">(.*?)&nbsp;-&nbsp;</span>(.*?)</div>`

	page_max = 750
)

type parseFunc func([]byte) model.ParseResult

type BaiduParser struct {
	ParseFunc parseFunc
	Name      string
}

func (p *BaiduParser) Parse(body []byte) (parseResult model.ParseResult) {
	return p.ParseFunc(body)
}

func (p *BaiduParser) Serialize() (name string, args []interface{}) {
	return p.Name, nil
}

func ParseBaidu(body []byte) (parseResult model.ParseResult) {
	var baiduItem baidumodel.BaiduItem

	re := regexp.MustCompile(re_baidu)
	match := re.FindAllSubmatch(body, -1)

	// logger.Logger.Infof("re_baidu: %v\n ", re_baidu)
	// logger.Logger.Infof("body: %v\n", string(body))
	for _, m := range match {
		baiduItem.Url = string(m[1])
		baiduItem.Title = string(m[2])
		baiduItem.UpdateTime = string(m[3])
		baiduItem.Intro = string(m[4])

		parseResult.Items = append(parseResult.Items, baiduItem)
	}
	return
}

func ParseBaiduBySelenium(body []byte) (parseResult model.ParseResult) {
	var baiduItem baidumodel.BaiduItem

	re := regexp.MustCompile(re_baidu_selenium)
	match := re.FindAllSubmatch(body, -1)

	// logger.Logger.Infof("re_baidu: %v\n ", re_baidu)
	// logger.Logger.Infof("body: %v\n", string(body))
	for _, m := range match {
		baiduItem.Url = string(m[1])
		baiduItem.Title = string(m[2])
		baiduItem.UpdateTime = string(m[3])
		baiduItem.Intro = string(m[4])

		baiduItem.Title = strings.Replace(baiduItem.Title, "<em>", "", -1)
		baiduItem.Title = strings.Replace(baiduItem.Title, "</em>", "", -1)
		baiduItem.Intro = strings.Replace(baiduItem.Intro, "<em>", "", -1)
		baiduItem.Intro = strings.Replace(baiduItem.Intro, "</em>", "", -1)
		parseResult.Items = append(parseResult.Items, baiduItem)
	}
	return
}
