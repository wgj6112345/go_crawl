package file

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/model/baidu"
	"testing"
)

func TestLetter(t *testing.T) {

	// fmt.Println("a" + "1")
	// fmt.Println(strconv.Itoa("146"))
}

func TestDiv(t *testing.T) {
	fmt.Println("div(2): ", div(2))
}

func TestSaveExcel(t *testing.T) {
	var (
		item1 baidu.BaiduItem
		item2 baidu.BaiduItem
	)
	InitExcel()

	item1 = baidu.BaiduItem{
		Title: "加快推进市域社会治理现代化(治理之道)--观点--人民网",
		Intro: "改革开放以来,随着经济社会不断发展,我国城镇化取得巨大成就,人口等各类要素越来越向市域聚集,市域在国家治理中的地位和作用日益凸显。市域层面具有较...",
		Url:   "baidu.com",
		// UpdateTime: time.Now(),
	}

	item2 = baidu.BaiduItem{
		Title: "“市域社会治理现代化”是什么?如何推进?",
		Intro: "8月26日上午,绍兴市市域社会治理现代化专题培训班在杭开班,区委政法委组织各镇街分管领导参加,培训期为三天。绍兴市委政法委常务副书记王荣彪作开班动...",
		Url:   "baidu.com",
		// UpdateTime: time.Now(),
	}

	SaveToExcel(item1)
	SaveToExcel(item2)

	SaveExcel("baidu")
}
