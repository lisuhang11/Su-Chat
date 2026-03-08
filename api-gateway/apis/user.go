package apis

import (
	"Su-chat/pkg/tools"
	"github.com/gin-gonic/gin"
)

// 用户注册
func Register(ctx *gin.Context) {
	tools.SuccessHttpResp(ctx, "测试用，先返回成功")
}

// 查询用户信息
func QueryUserInfo(ctx *gin.Context) {
	tools.SuccessHttpResp(ctx, "测试用，先返回成功")
}

// 更新用户信息
func UpdateUserInfo(ctx *gin.Context) {
	tools.SuccessHttpResp(ctx, "测试用，先返回成功")
}

// 查询在线状态
func QueryUserOnlineStatus(ctx *gin.Context) {
	tools.SuccessHttpResp(ctx, "测试用，先返回成功")
}
