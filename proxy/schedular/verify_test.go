package schedular

import (
	"fmt"
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	"net/http"
	"net/http/httputil"
	"regexp"
	"testing"
)

func TestVerify(t *testing.T) {
	// t.Skip()
	proxy := "40.79.26.139:1080"
	ip := model.IP{
		Ip: proxy,
	}
	verify(ip)
}

func TestProxy(t *testing.T) {
	t.Skip()
	resp, err := http.Get("http://127.0.0.1:9191/ip")
	if err != nil {
		logger.Logger.Errorf("get proxy failed, err : %v\n", err)
		return
	}
	defer resp.Body.Close()

	reProxy := `([\d]+\.[\d]+\.[\d]+\.[\d]+:[\d]+)`
	re := regexp.MustCompile(reProxy)
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		logger.Logger.Errorf("DumpResponse failed, err : %v\n", err)
		return
	}

	fmt.Println(string(body))
	match := re.FindString(string(body))

	fmt.Println("match: ", match)

}
