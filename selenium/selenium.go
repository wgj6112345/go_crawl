package selenium

import (
	"errors"
	"fmt"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Crawler struct {
	ChromeDriver string
	Port         int
	Service      *selenium.Service
	Caps         selenium.Capabilities
}

//开启驱动服务
func NewCrawler() (*Crawler, error) {
	crawler := &Crawler{
		ChromeDriver: `C:/Users/wgj61/Downloads/chromedriver_win32/chromedriver.exe`, //google浏览器驱动
		Port:         9515,
		Service:      nil,
	}
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(crawler.ChromeDriver, crawler.Port, opts...)
	if nil != err {
		return nil, errors.New("start a chromedriver service falid," + err.Error())
	}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2, //不加载图片，提高浏览器响应速度
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", //不弹出窗口
			"--user-agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", // 模拟user-agent，防反爬
		},
	}
	//以上是设置浏览器参数
	caps.AddChrome(chromeCaps)
	crawler.Service = service
	crawler.Caps = caps
	return crawler, nil
}

//打开窗口
func (c *Crawler) NewRemote() (selenium.WebDriver, error) {
	w_b1, err := selenium.NewRemote(c.Caps, fmt.Sprintf("http://localhost:%d/wd/hub", c.Port))
	if err != nil {
		return nil, errors.New("connect to the webDriver faild," + err.Error())
	}
	return w_b1, nil
}

//关闭驱动服务
func (c *Crawler) Shutdown() {
	_ = c.Service.Stop()
}

func (c *Crawler) Fetch(url string) (body []byte, err error) {
	var (
		driver selenium.WebDriver
		resp   string
		// cookie *selenium.Cookie
	)

	if driver, err = c.NewRemote(); err != nil {
		err = errors.New("connect to the webDriver faild," + err.Error())
	}
	defer driver.Quit()

	// cookie = &selenium.Cookie{Name: "BDUSS", Value: "00WGhPbldDOXp6MVBmZ1NzVG1MQkxDeTRxZ1BoN2R-Zn5SVnV3czNDVTNoaHRmRVFBQUFBJCQAAAAAAAAAAAEAAADbgMcdd2dqNjExMjM0NQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAADf58143-fNeT0"}
	// if err = driver.AddCookie(cookie); err != nil {
	// 	err = errors.New("Add cookie failed, " + err.Error())
	// }
	// <-time.Tick(time.Duration(time.Second * 2))

	if err = driver.Get(url); err != nil {
		err = errors.New("webDriver get url faild," + err.Error())
	}

	if resp, err = driver.PageSource(); err != nil {
		err = errors.New("webDriver get pagesource faild," + err.Error())
	}

	body = []byte(resp)

	// time.Sleep(time.Second * 1)
	// <-time.Tick(time.Duration(time.Second * 3))
	return
}
