package schedular

import (
	"github.com/wgj6112345/go_crawl/proxy/logger"
	"github.com/wgj6112345/go_crawl/proxy/model"
	"net/http"
	"net/url"
	"time"
)

var (
	responseTime float64 = 3.0
)

func verify(ipProxy model.IP) (isAvailable bool) {
	timeout := time.Duration(time.Second * 3)
	Url := "http://httpbin.org/get"
	// Url := "https://book.douban.com/tag/"

	req, _ := http.NewRequest("GET", Url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")

	proxy, _ := url.Parse(ipProxy.Ip)
	// if err != nil {
	// 	logger.Logger.Errorf("url parse proxy: %v failed, err : %v\n", ipProxy.Ip, err)
	// 	return
	// }
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: timeout,
	}

	start := time.Now().UnixNano()
	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorf("client get url: %v failed\n", Url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logger.Logger.Debugf("url: %v is not available\n", Url)
		return
	}

	end := time.Now().UnixNano()
	spend := float64(end-start) / 1000000000.0

	if spend > responseTime {
		logger.Logger.Errorf("url: %v is too slow spend: %v\n", Url, spend)
		return
	}

	logger.Logger.Infof("ip: %v is available, spend: %v\n", ipProxy.Ip, spend)
	return true
}
