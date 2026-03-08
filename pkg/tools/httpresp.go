package tools

import (
	"Su-chat/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SuccessHttpResp 返回成功响应
// httpCode: HTTP 状态码（如 200），data: 返回的数据
func SuccessHttpResp(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

// ErrorHttpResp 返回错误响应
// httpCode: HTTP 状态码（如 400），bizCode: 业务错误码
func ErrorHttpResp(ctx *gin.Context, httpCode int, bizCode errs.IMErrorCode) {
	ctx.JSON(httpCode, gin.H{
		"code": int(httpCode),
		"msg":  errs.GetErrorInfoByCode(bizCode),
		"data": nil,
	})
}
