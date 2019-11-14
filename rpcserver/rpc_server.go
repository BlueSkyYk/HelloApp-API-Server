package rpcserver

import (
	"HelloApp_Api/rpcprotos"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

/**
	初始化RPC服务
 */
func StartRpcServer(port int16) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("启动Rpc服务失败 : %v", err)
	}
	fmt.Println("启动Rpc服务成功")
	grpcServer := grpc.NewServer()
	rpcprotos.RegisterGoRpcServerInterfaceServer(grpcServer, &GoRpcServerInterfaceServerImpl{})
	grpcServer.Serve(lis)
}
