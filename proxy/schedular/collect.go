package schedular

import (
	. "github.com/wgj6112345/go_crawl/proxy/collector"
	"github.com/wgj6112345/go_crawl/proxy/model"
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
