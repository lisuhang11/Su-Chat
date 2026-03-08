package router

import (
	"Su-chat/api-gateway/apis"
	"github.com/gin-gonic/gin"
)

var GE *gin.Engine

func InitRouter() *gin.Engine {
	GE = gin.Default() // 初始化引擎

	// 注册路由
	GE.POST("/users/register", apis.Register)
	GE.GET("/users/info", apis.QueryUserInfo)
	GE.POST("/users/onlinestatus/query", apis.QueryUserOnlineStatus)
	GE.POST("/users/update", apis.UpdateUserInfo)
	return GE
}
