# go_crawl


## 文件目录结构
```
├─conf  
├─db
├─doc
├─engine  // 爬虫引擎
├─fetcher // 获取 url 对应的 内容
├─file    // excel文件输出目录
├─logger  // 日志组件
├─logs    // 日志存放目录
├─model  // 
│  ├─baidu
│  └─book
├─parser  // 网页内容解析
│  └─baidu
├─proxy   // 网络代理
│  ├─collector // 代理搜集
│  ├─doc
│  ├─logger
│  ├─logs
│  ├─main
│  │  ├─collector 
│  │  └─server 
│  ├─model 
│  ├─redis
│  ├─schedular 
│  └─tools
├─rpc
│  ├─config
│  ├─ElasticService
│  │  ├─client
│  │  ├─server
│  │  └─service
│  ├─rpcHelper
│  └─WorkerService
│      ├─client
│      ├─server
│      └─service
├─schedular
└─selenium 

```


## 使用方法

```
# 如果访问的网站反爬很严格，可以先启动 网络代理服务器
# 因为要用到 redis 存放代理池，同时需要启动redis服务
cd ./go_crawl/proxy/main/server
go build
./server.exe


# 如果不要代理直接启动 爬虫服务
cd ./go_crawl
go build
./go_crawl.exe

tips: 很不幸的是，我试验的两个网站 豆瓣 和 百度，爬了一两百条之后，直接不能访问了，后面尝试用 selenium 也不奏效，有点蛋疼。
```