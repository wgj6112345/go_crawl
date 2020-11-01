package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/wgj6112345/go_crawl/engine"
	"github.com/wgj6112345/go_crawl/model"
	"github.com/wgj6112345/go_crawl/parser"
	"github.com/wgj6112345/go_crawl/schedular"
)

func main() {

	// 性能分析
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	time.Sleep(20 * time.Second)

	var (
		url = "https://book.douban.com/tag/"
		// url = "https://book.douban.com/tag/%E7%A5%9E%E7%BB%8F%E7%BD%91%E7%BB%9C"
		// url = "https://book.douban.com/subject/30293801/"
	)

	s := schedular.QueueSchedular{}
	e := engine.ConCurrentEngine{
		Schedular:   &s,
		WorkNum:     10,
		ItemChan:    schedular.SaveItem(),
		ProcessFunc: engine.Work,
	}

	parseInst := &parser.Level12Parser{}
	parseInst.ParseFunc = parser.ParseLevel1

	go e.Run(model.Request{
		Url:    url,
		Parser: parseInst,
	})

	time.Sleep(30 * time.Second)
	// 内存分析
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}

		// runtime.GC()
		pprof.WriteHeapProfile(file)
		file.Close()
	}

	// var (

	// 	baiduBase          = "https://www.baidu.com/s"
	// 	keyWord     string = "市域社会治理现代化"
	// 	maxPage     int    = 75
	// 	requestList []model.Request
	// )

	// 构造 url 集合  https://www.baidu.com/s?wd=市域社会治理现代化&pn=710
	// for page := 0; page <= maxPage; page++ {
	// 	v := url.Values{}
	// 	v.Set("wd", keyWord)
	// 	v.Set("pn", fmt.Sprintf("%v", page*10))
	// 	urlEncode := v.Encode()
	// 	Url := fmt.Sprintf("%s?%s", baiduBase, urlEncode)
	// 	requestList = append(requestList, model.Request{
	// 		Url: Url,
	// 		Parser: &baidu.BaiduParser{
	// 			ParseFunc: baidu.ParseBaiduBySelenium,
	// 			Name:      "ParseBaiduBySelenium",
	// 		},
	// 	})
	// }
	// s := schedular.QueueSchedular{}
	// e := engine.ConCurrentEngine{
	// 	Schedular:   &s,
	// 	WorkNum:     60,
	// 	ItemChan:    schedular.SaveBaiduItem(),
	// 	ProcessFunc: engine.WorkBySelenium,
	// }

	// e.Run(requestList...)
}
