package collector

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/proxy/logger"
	"github.com/wgj6112345/go_crawl/proxy/model"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	url_data5u = "http://www.data5u.com/" // ?page= &country="中国"

)

var (
	re_data5u = `<ul class="l2">[\d\D]*?<span><li>([^<]+)</li></span>[\d\D]*?<span[\d\D]*?><li class="port HZZZC">([^<]+)</li></span>[\d\D]*?<span[\d\D]*?><li>.*?</li></span>[\d\D]*?<span[\d\D]*?><li>([^<]+)</li></span>[\d\D]*?<span><li>.*?</li></span>[\d\D]*?<span[\d\D]*?><li>([^<]+)</li></span>[\d\D]*?<span[\d\D]*?><li>.*?</li></span>[\d\D]*?<span[\d\D]*?><li>([^<]+)</li></span>[\d\D]*?<span[\d\D]*?><li>.*?</li></span>[\d\D]*?<div[\d\D]*?></div>[\d\D]*?</ul>`
)

func GetData5u(wg *sync.WaitGroup, ipChan chan model.IP) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		// logger.Logger.Debugf("fetch Url: %v\n", url_data5u)
		body, err := Fetch(url_data5u)
		if err != nil {
			logger.Logger.Errorf("fetch url: %v failed, err : %v\n", url_data5u, err)
			return
		}
		ipList := parseData5u(body)
		for _, ipItem := range ipList {
			ipChan <- ipItem
		}

	}()

}

func parseData5u(body []byte) (ipList []model.IP) {
	re := regexp.MustCompile(re_data5u)
	match := re.FindAllSubmatch(body, -1)

	var ip model.IP
	for _, m := range match {

		ip.Ip = fmt.Sprintf("%s:%s", string(m[1]), string(m[2]))

		ip.Type = string(m[3])
		ip.Location = string(m[4])
		// 判断毫秒 还是 秒
		var speed float64

		speedStr := string(m[5])
		// fmt.Println("speedStr: ", speedStr)
		if strings.Contains(speedStr, "毫") {
			speedSplits := strings.Split(speedStr, "毫秒")
			speedstr := speedSplits[0]
			speed, _ = strconv.ParseFloat(speedstr, 64)
			speed = speed / 1000
		} else {
			speedSplits := strings.Split(speedStr, "秒")
			speedstr := speedSplits[0]
			speed, _ = strconv.ParseFloat(speedstr, 64)
		}
		ip.Speed = speed

		ipList = append(ipList, ip)
	}

	return
}
