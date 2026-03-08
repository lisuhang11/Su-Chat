package main

import (
	"Su-chat/services/user/gen"
	"Su-chat/services/user/internal/logic"
	"Su-chat/services/user/internal/storages"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// 初始化数据库连接
	db, err := storages.InitDB()
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 1. 监听端口
	listen, err := net.Listen("tcp", ":10001")
	if err != nil {
		panic(fmt.Sprintf("监听端口失败: %v", err))
	}
	fmt.Println("User gRPC 服务端启动，监听端口 10001...")

	// 2. 创建gRPC服务器实例
	grpcServer := grpc.NewServer()

	// 3. 注册用户服务实现
	gen.RegisterUserServiceServer(grpcServer, logic.NewUserServer(db))

	// 4. 启动服务器
	if err := grpcServer.Serve(listen); err != nil {
		panic(fmt.Sprintf("服务器启动失败: %v", err))
	}
}
