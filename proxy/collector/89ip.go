package collector

import (
	"github.com/wgj6112345/go_crawl/proxy/logger"
	"github.com/wgj6112345/go_crawl/proxy/model"
	"regexp"
	"strings"
	"sync"
)

const (
	url_89ip = "http://www.89ip.cn/tqdl.html?api=1&num=100&port=&address=%E7%BE%8E%E5%9B%BD&isp="
)

var (
	re_89ip = `([\d]{2,3}.*?)<br`
)

func Get89Ip(wg *sync.WaitGroup, ipChan chan model.IP) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		body, err := Fetch(url_89ip)
		if err != nil {
			logger.Logger.Errorf("fetch url: %v failed, err : %v\n", url_89ip, err)
			return
		}
		ipList := parse89Ip(body)
		for _, ipItem := range ipList {
			ipChan <- ipItem
		}
	}()
 
}

func parse89Ip(body []byte) (ipList []model.IP) {
	re := regexp.MustCompile(re_89ip)
	match := re.FindAllSubmatch(body, -1)

	var ip model.IP
	for index, m := range match {
		if index == 0 {
			continue
		}

		ip.Ip = strings.TrimSpace(string(m[1]))
		ipList = append(ipList, ip)
	}

	return
}
