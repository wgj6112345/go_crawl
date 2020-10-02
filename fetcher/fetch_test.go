package fetcher

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/parser/baidu"
	"testing"
)

var (
	url1 = `https://www.baidu.com/s?pn=110&wd=%E5%B8%82%E5%9F%9F%E7%A4%BE%E4%BC%9A%E6%B2%BB%E7%90%86%E7%8E%B0%E4%BB%A3%E5%8C%96`
)

func TestFetchBaidu(t *testing.T) {

	body, _ := FetchByProxy(url1)
	fmt.Println("body: ", string(body))
	result := baidu.ParseBaidu(body)
	fmt.Println("result: ", result)
}
