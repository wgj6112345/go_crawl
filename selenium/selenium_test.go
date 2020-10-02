package selenium

import (
	"fmt"
	"testing"
)

var url = "https://www.baidu.com/s?wd=%E5%B8%82%E5%9F%9F%E7%A4%BE%E4%BC%9A%E6%B2%BB%E7%90%86%E7%8E%B0%E4%BB%A3%E5%8C%96&pn=10&oq=%E5%B8%82%E5%9F%9F%E7%A4%BE%E4%BC%9A%E6%B2%BB%E7%90%86%E7%8E%B0%E4%BB%A3%E5%8C%96"

func TestSelenium(t *testing.T) {
	crawler, _ := NewCrawler()
	driver, _ := crawler.NewRemote()

	driver.Get(url)
	driver.SetImplicitWaitTimeout(3)
	body, _ := driver.PageSource()

	fmt.Println("body: ", body)
	defer driver.Quit()
	// defer crawler.Service.Stop()
}
