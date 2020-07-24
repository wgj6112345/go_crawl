package collector

import (
	"fmt"
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	"imooc/分布式爬虫项目/demo1/proxy/tools"
	"net/url"
	"regexp"
	"sync"
	"time"
)

const (
	url_Jxl = "https://ip.jiangxianli.com/" // ?page= &country="中国"
)

var (
	re_Jxl = `<tr><td>([^<]+)</td><td>([^<]+)</td><td>.*?</td><td>.*?</td><td>([^<]+)</td><td>.*?</td><td>.*?</td><td>([^<]+)</td><td>.*?</td><td>([^<]+)</td><td><button[\d\D]*?</button></td>`
)

func GetJxl(wg *sync.WaitGroup, ipChan chan model.IP) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		var ipList []model.IP
		for page := 1; page < 7; page++ {
			v := url.Values{}
			v.Set("page", fmt.Sprintf("%v", page))
			v.Set("country", "中国")
			urlEncode := v.Encode()
			Url := fmt.Sprintf("%s?%s", url_Jxl, urlEncode)

			// logger.Logger.Debugf("fetch Url: %v\n", Url)
			body, err := Fetch(Url)
			if err != nil {
				logger.Logger.Errorf("fetch url: %v failed, err : %v\n", url_Jxl, err)
				return
			}
			ips := parseJxl(body)
			ipList = append(ipList, ips...)
			time.Sleep(time.Millisecond * 500)
		}

		for _, ipItem := range ipList {
			ipChan <- ipItem
		}

	}()

}

func parseJxl(body []byte) (ipList []model.IP) {
	re := regexp.MustCompile(re_Jxl)
	match := re.FindAllSubmatch(body, -1)

	var ip model.IP
	for _, m := range match {
		ip.Ip = fmt.Sprintf("%s:%s", string(m[1]), string(m[2]))
		// ip.Location = string(m[3])

		// 判断毫秒 还是 秒
		// var speed float64

		// speedStr := string(m[4])
		// fmt.Println("speedStr: ", speedStr)
		// if strings.Contains(speedStr, "毫") {
		// 	speedSplits := strings.Split(speedStr, "毫秒")
		// 	speedstr := speedSplits[0]
		// 	speed, _ = strconv.ParseFloat(speedstr, 64)
		// 	speed = speed / 1000
		// } else {
		// 	speedSplits := strings.Split(speedStr, "秒")
		// 	speedstr := speedSplits[0]
		// 	speed, _ = strconv.ParseFloat(speedstr, 64)
		// }
		// ip.Speed = speed
		ip.LastVerify = tools.TimeConvertWithBar(string(m[5]))

		ipList = append(ipList, ip)
	}

	return
}
