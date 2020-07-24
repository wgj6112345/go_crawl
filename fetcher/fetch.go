package fetcher

import (
	"bufio"
	"imooc/分布式爬虫项目/demo1/logger"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
	"time"

	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// var (
// 	// 错误次数 连续错误超过 10次  fetch 暂定 几秒
// 	errorCount = 0
// )

func FetchByProxy(Url string) (body []byte, err error) {
	// redis 布隆过滤器 判断 URL 是否 已经访问过

	var timeout = time.Duration(3 * time.Second)
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		logger.Logger.Errorf("http.NewRequest url: %v failed, err : %v\n", Url, err)
		return
	}
	// SetFakeHeader(req)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")

	proxyAddr := getProxy()
	if proxyAddr == "" {
		// 暂定几秒
		time.Sleep(time.Second * 3)
	}

	proxy, _ := url.Parse(proxyAddr)
	// if err != nil {
	// 	logger.Logger.Errorf("url parse proxy: %v failed, err : %v\n", proxyAddr, err)
	// 	return FetchByProxy(Url)
	// }

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Logger.Errorf("get url: %v failed, err : %v\n", Url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logger.Logger.Errorf("use proxy: %v fetch url: %v is not available\n", proxyAddr, Url)
		return
	}

	bufReader := bufio.NewReader(resp.Body)
	encode := checkEncoding(bufReader)
	utf8Reader := transform.NewReader(bufReader, encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func checkEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		logger.Logger.Errorf("peek err : %v\n", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func getProxy() (match string) {
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

	match = re.FindString(string(body))

	match = strings.TrimSpace(match)
	return match
}

func Fetch(url string) ([]byte, error) {
	// 控制爬虫频率
	<-time.Tick(time.Duration(time.Second))

	// TODO:  用 布隆过滤器 进行过滤 已经访问过的网页 就不再访问了

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")

	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		logger.Logger.Errorf("Error status code : %v\n", resp.StatusCode)
	}
	defer resp.Body.Close()

	bufReader := bufio.NewReader(resp.Body)
	encode := checkEncoding(bufReader)
	utf8Reader := transform.NewReader(bufReader, encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}
