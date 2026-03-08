package main

import (
	"Su-chat/services/auth/gen"
	"Su-chat/services/auth/internal/logic"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {

	// 1. 监听端口
	listen, err := net.Listen("tcp", ":10002")
	if err != nil {
		panic(fmt.Sprintf("监听端口失败: %v", err))
	}
	fmt.Println("Auth gRPC 服务端启动，监听端口 10002...")

	// 2. 创建gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 3. 注册用户服务实现
	gen.RegisterAuthServiceServer(grpcServer, logic.NewAuthServer())

	// 4. 启动服务器
	if err := grpcServer.Serve(listen); err != nil {
		panic(fmt.Sprintf("服务器启动失败: %v", err))
	}

}
