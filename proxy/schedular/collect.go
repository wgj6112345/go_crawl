package schedular

import (
	. "imooc/分布式爬虫项目/demo1/proxy/collector"
	"imooc/分布式爬虫项目/demo1/proxy/model"
	"sync"
)

func Collect(ipChan chan model.IP) {
	var wg sync.WaitGroup
	Get66Ip(&wg, ipChan)
	Get89Ip(&wg, ipChan)
	GetData5u(&wg, ipChan)
	Get3366Ip(&wg, ipChan)
	GetJxl(&wg, ipChan)
	GetKXDL(&wg, ipChan)

	wg.Wait()
}
