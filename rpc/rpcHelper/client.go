package rpcHelper

import (
	"imooc/分布式爬虫项目/demo1/logger"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func NewClient(host string) (client *rpc.Client) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		logger.Logger.Errorf("dial host: %v failed, err: %v\n", host, err)
		return
	}
	client = rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return
}
