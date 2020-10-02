package rpcHelper

import (
	"github.com/wgj6112345/go_crawl/logger"
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
