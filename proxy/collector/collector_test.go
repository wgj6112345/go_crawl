package collector

import (
	"fmt"
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	"regexp"
	"sync"
	"testing"
)

// const (
// 	url_66ip          = "//www.66ip.cn/mo.php?tqsl=100"
// 	url_89ip          = "//www.89ip.cn/tqdl.html?api=1&num=100&port=&address=%E7%BE%8E%E5%9B%BD&isp="
// 	url_ip3366        = "//www.ip3366.net/free/?stype=1&page=1"
// 	url_kuaidl        = "//www.kuaidaili.com/"
// 	url_proxylist     = "//list.proxylistplus.com/Fresh-HTTP-Proxy-List-1"
// 	url_proxylist_ssl = "//list.proxylistplus.com/SSL-List-1"
// 	url_jiangxianli   = "//ip.jiangxianli.com/?page=1"
// 	url_kaixindl      = "//www.kxdaili.com/dailiip.html"
// 	url_data5u        = "//www.data5u.com"
// )

func TestSourceUrl(t *testing.T) {
	t.Skip()
	urlList := []string{
		"http://www.66ip.cn/mo.php?tqsl=100",
		"http://www.89ip.cn/tqdl.html?api=1&num=100&port=&address=%E7%BE%8E%E5%9B%BD&isp=",
		"http://www.ip3366.net/free/?stype=1&page=1",
		" https://www.kuaidaili.com/",
		"https://list.proxylistplus.com/Fresh-HTTP-Proxy-List-1",
		"https://list.proxylistplus.com/SSL-List-1",
		" https://ip.jiangxianli.com/?page=1",
		"http://www.kxdaili.com/dailiip.html",
		"http://www.data5u.com",
	}

	for _, url := range urlList {
		fmt.Printf("fetch url: %v...\n", url)
		_, err := Fetch(url)
		if err != nil {
			logger.Logger.Errorf("get url: %v failed, err : %v\n", url, err)
			return
		}
	}

}

func TestOneUrl(t *testing.T) {
	t.Skip()
	url := "http://www.66ip.cn/mo.php?tqsl=100"
	body, err := Fetch(url)
	if err != nil {
		logger.Logger.Errorf("get url: %v failed, err : %v\n", url, err)
		return
	}

	fmt.Println("body: ", string(body))

	ipList := parse66Ip(body)
	fmt.Println("ipList: ", ipList)
}

func TestGetIp(t *testing.T) {
	t.Skip()
	// Get66Ip()
	// for {
	// 	select {
	// 	case ip := <-ipChan:
	// 		fmt.Println(ip)
	// 	default:
	// 		time.Sleep(time.Millisecond * 50)
	// 	}

	// }
}

func TestRegexp(t *testing.T) {
	t.Skip()
	str := `            <tr>
                <td>45.80.104.45</td>
                <td>8085</td>
                <td>高匿代理IP</td>
                <td>HTTP</td>
                <td>高匿_欧盟</td>
                <td>1秒</td>
                <td>2020/7/22 11:10:05</td>
            </tr>
           
            <tr>
                <td>45.80.104.68</td>
                <td>8085</td>
                <td>高匿代理IP</td>
                <td>HTTPS</td>
                <td>SSL高匿_欧盟</td>
                <td>4秒</td>
                <td>2020/7/22 11:10:04</td>
            </tr>
           
            <tr>
                <td>45.80.104.55</td>
                <td>8085</td>
                <td>高匿代理IP</td>
                <td>HTTP</td>
                <td>高匿_欧盟</td>
                <td>10秒</td>
                <td>2020/7/22 10:39:11</td>
            </tr>`

	// test, _ := Fetch(url_3366ip)
	restr := `<tr>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>.*?</td>[\d\D]*?<td>.*?</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?<td>([^<]+)秒</td>[\d\D]*?<td>([^<]+)</td>[\d\D]*?</tr>`

	// restr := `<tr>[\d\D]*?<td>[\d\D]*?</td>[\d\D]*?</tr>`
	re := regexp.MustCompile(restr)

	// fmt.Println("test: ", string(test))
	match := re.FindAllSubmatch([]byte(str), -1)
	for _, m := range match {
		// fmt.Println("m[0]: ", string(m[0]))
		fmt.Println("m[1]: ", string(m[1]))
		fmt.Println("m[2]: ", string(m[2]))
		fmt.Println("m[3]: ", string(m[3]))
		fmt.Println("m[4]: ", string(m[4]))
		fmt.Println("m[5]: ", string(m[5]))
	}
}

func TestRun(t *testing.T) {
	ipChan := make(chan model.IP, 1000)

	var wg sync.WaitGroup
	Get66Ip(&wg, ipChan)
	Get89Ip(&wg, ipChan)
	GetData5u(&wg, ipChan)
	Get3366Ip(&wg, ipChan)
	GetJxl(&wg, ipChan)
	GetKXDL(&wg, ipChan)

	wg.Wait()
	// schedular.Collect(ipChan)

	// var count = 0
	// for ip := range ipChan {
	// 	// count++
	// 	fmt.Println(ip)
	// 	// fmt.Println("count: ", count)
	// }

}
