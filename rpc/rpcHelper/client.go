package rpcHelper

import (
	"github.com/wgj6112345/go_crawl/logger"
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
