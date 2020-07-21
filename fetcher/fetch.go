package fetcher

import (
	"bufio"
	"imooc/分布式爬虫项目/demo1/logger"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/transform"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string) ([]byte, error) {
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

func checkEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		logger.Logger.Errorf("peek err : %v\n", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
