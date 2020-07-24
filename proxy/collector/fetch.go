package collector

import (
	"bufio"
	"imooc/分布式爬虫项目/demo1/proxy/logger"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var userAgent = [...]string{"Mozilla/5.0 (compatible, MSIE 10.0, Windows NT, DigExt)",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, 360SE)",
	"Mozilla/4.0 (compatible, MSIE 8.0, Windows NT 6.0, Trident/4.0)",
	"Mozilla/5.0 (compatible, MSIE 9.0, Windows NT 6.1, Trident/5.0,",
	"Opera/9.80 (Windows NT 6.1, U, en) Presto/2.8.131 Version/11.11",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, TencentTraveler 4.0)",
	"Mozilla/5.0 (Windows, U, Windows NT 6.1, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Macintosh, Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11", "Mozilla/5.0 (Macintosh, U, Intel Mac OS X 10_6_8, en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	"Mozilla/5.0 (Linux, U, Android 3.0, en-us, Xoom Build/HRI39) AppleWebKit/534.13 (KHTML, like Gecko) Version/4.0 Safari/534.13",
	"Mozilla/5.0 (iPad, U, CPU OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"Mozilla/4.0 (compatible, MSIE 7.0, Windows NT 5.1, Trident/4.0, SE 2.X MetaSr 1.0, SE 2.X MetaSr 1.0, .NET CLR 2.0.50727, SE 2.X MetaSr 1.0)",
	"Mozilla/5.0 (iPhone, U, CPU iPhone OS 4_3_3 like Mac OS X, en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
	"MQQBrowser/26 Mozilla/5.0 (Linux, U, Android 2.3.7, zh-cn, MB200 Build/GRJ22, CyanogenMod-7) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
}

func Fetch(url string) ([]byte, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")

	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		logger.Logger.Errorf("Error status code : %v\n", resp.StatusCode)
		return Fetch(url)
	}
	defer resp.Body.Close()

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

func FetchByProxy(proxyAddr string, Url string) (body []byte, err error) {
	var timeout = time.Duration(3 * time.Second)
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		logger.Logger.Errorf("http.NewRequest url: %v failed, err : %v\n", Url, err)
		return
	}
	SetFakeHeader(req)

	proxy, err := url.Parse(proxyAddr)
	if err != nil {
		logger.Logger.Errorf("url parse proxy: %v failed, err : %v\n", proxyAddr, err)
		return
	}

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

	bufReader := bufio.NewReader(resp.Body)
	encode := checkEncoding(bufReader)
	utf8Reader := transform.NewReader(bufReader, encode.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func getRandomUserAgent() string {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	random := r.Intn(len(userAgent))
	return userAgent[random]
}

func SetFakeHeader(req *http.Request) {
	req.Header.Set("User-Agent", getRandomUserAgent())
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Connection", "keep-alive")
}
