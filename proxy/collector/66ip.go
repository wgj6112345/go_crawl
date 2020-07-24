package collector

import (
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	"regexp"
	"strings"
	"sync"
)

const (
	url_66ip = "http://www.66ip.cn/mo.php?tqsl=100"
)

var (
	re_66ip = `(.*?)<br`
)

func Get66Ip(wg *sync.WaitGroup, ipChan chan model.IP) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		body, err := Fetch(url_66ip)
		if err != nil {
			logger.Logger.Errorf("fetch url: %v failed, err : %v\n", url_66ip, err)
			return
		}

		ipList := parse66Ip(body)
		for _, ipItem := range ipList {
			ipChan <- ipItem
		}

	}()
}

func parse66Ip(body []byte) (ipList []model.IP) {
	re := regexp.MustCompile(re_66ip)
	match := re.FindAllSubmatch(body, -1)

	var ip model.IP

	for _, m := range match {
		ip.Ip = strings.TrimSpace(string(m[1]))
		ipList = append(ipList, ip)
	}

	return
}
