package collector

import (
	"fmt"
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	url_KXDL = "http://www.kxdaili.com/dailiip/" // ?page= &country="中国"

)

var (
	re_KXDL = `<tr>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>.*?</td>[\d\D]*?<td>.*?</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>.*?</td>[\d\D]*?</tr>`
)

func GetKXDL(wg *sync.WaitGroup, ipChan chan model.IP) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		var ipList []model.IP
		for cate := 1; cate < 3; cate++ {
			for page := 1; page < 8; page++ {
				Url := fmt.Sprintf("%s%v/%v.html/", url_KXDL, cate, page)

				// logger.Logger.Debugf("fetch Url: %v\n", Url)
				body, err := Fetch(Url)
				if err != nil {
					logger.Logger.Errorf("fetch url: %v failed, err : %v\n", url_KXDL, err)
					return
				}
				ips := parseKXDL(body)
				ipList = append(ipList, ips...)
				time.Sleep(time.Millisecond * 500)
			}
		}
		for _, ipItem := range ipList {
			ipChan <- ipItem
		}

	}()

}

func parseKXDL(body []byte) (ipList []model.IP) {
	re := regexp.MustCompile(re_KXDL)
	match := re.FindAllSubmatch(body, -1)

	var ip model.IP
	for _, m := range match {

		ip.Ip = fmt.Sprintf("%s:%s", string(m[1]), string(m[2]))
		ip.Location = string(m[4])

		// 判断毫秒 还是 秒
		var speed float64

		speedStr := string(m[3])
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
