package rpcHelper

import (
	"imooc/分布式爬虫项目/demo1/logger"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRPC(host string, service interface{}) {
	rpc.Register(service)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		logger.Logger.Errorf("listen tcp failed, err: %v\n", err)
		return
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			logger.Logger.Errorf("establish connect failed, err: %v\n", err)
			return
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
