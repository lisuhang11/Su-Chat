package main

import (
	"Su-chat/api-gateway/internal/router"
)

func main() {
	r := router.InitRouter() // 初始化路由
	r.Run(":8080")           // 启动服务，监听 8080 端口
}
