package collector

import (
	"fmt"
	"github.com/wgj6112345/go_crawl/proxy/logger"
	"github.com/wgj6112345/go_crawl/proxy/model"
	"github.com/wgj6112345/go_crawl/proxy/tools"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	url_KDL = "https://www.kuaidaili.com/free/inha/" // ?page= &country="中国"

)

var (
	re_KDL = `tr>[\d\D]*?<td data-title="IP">([^<]+)</td>[\d\D]*?<td data-title="PORT">([^<]+)</td>[\d\D]*?<td data-title="匿名度">.*?</td>[\d\D]*?<td data-title="类型">.*?</td>[\d\D]*?<td data-title="位置">([^<]+)</td>[\d\D]*?<td data-title="响应速度">([^<]+)</td>[\d\D]*?<td data-title="最后验证时间">([^<]+)</td>[\d\D]*?</tr>`
)

func GetKDL(wg *sync.WaitGroup, ipChan chan model.IP) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		var ipList []model.IP
		for page := 1; page < 20; page++ {
			Url := fmt.Sprintf("%s%v/", url_KDL, page)

			// logger.Logger.Debugf("fetch Url: %v\n", Url)
			body, err := Fetch(Url)
			if err != nil {
				logger.Logger.Errorf("fetch url: %v failed, err : %v\n", url_KDL, err)
				return
			}
			ips := parseKDL(body)
			ipList = append(ipList, ips...)
			time.Sleep(time.Second)
		}
		for _, ipItem := range ipList {
			ipChan <- ipItem
		}

	}()

}

func parseKDL(body []byte) (ipList []model.IP) {
	re := regexp.MustCompile(re_KDL)
	match := re.FindAllSubmatch(body, -1)

	var ip model.IP
	for _, m := range match {
		ip.Ip = fmt.Sprintf("%s:%s", string(m[1]), string(m[2]))
		ip.Location = string(m[3])

		// 判断毫秒 还是 秒
		var speed float64

		speedStr := string(m[4])
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
		ip.LastVerify = tools.TimeConvertWithBar(string(m[5]))

		ipList = append(ipList, ip)
	}

	return
}
