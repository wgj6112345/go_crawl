package collector

import (
	"fmt"
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	"imooc/分布式爬虫项目/demo1/proxy/tools"
	"regexp"
	"strconv"
	"sync"
)

const (
	url_3366ip = "http://www.ip3366.net/free/?stype=1&page=1"
)

var (
	re_3366ip = `<tr>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>.*?</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>([^<]+)秒</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?</tr>`
)

func Get3366Ip(wg *sync.WaitGroup, ipChan chan model.IP) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		body, err := Fetch(url_3366ip)
		if err != nil {
			logger.Logger.Errorf("fetch url: %v failed, err : %v\n", url_89ip, err)
			return
		}
		ipList := parse3366Ip(body)
		for _, ipItem := range ipList {
			ipChan <- ipItem
		}

	}()
 
}

func parse3366Ip(body []byte) (ipList []model.IP) {
	re := regexp.MustCompile(re_3366ip)
	match := re.FindAllSubmatch(body, -1)

	var ip model.IP
	for _, m := range match {
		ip.Ip = fmt.Sprintf("%s:%s", string(m[1]), string(m[2]))
		ip.Type = string(m[3])
		ip.Location = string(m[4])

		speed, _ := strconv.ParseFloat(string(m[5]), 64)
		ip.Speed = speed
		ip.LastVerify = tools.TimeConvertWithSlash(string(m[6]))

		ipList = append(ipList, ip)
	}

	return
}
