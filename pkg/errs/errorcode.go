package errs

/*
0 : success
10000~10999 : api
11000~11999 : connect
12000~12999 : private msg
13000~13999 : group
14000~14999 : chatroom
*/

type IMErrorCode int

// user errorcode
const (
	USER_NOT_EXIST IMErrorCode = 10102
)

// 错误码对应的错误信息
var ErrorInfo = map[IMErrorCode]string{
	USER_NOT_EXIST: "用户不存在",
}

func GetErrorInfoByCode(code IMErrorCode) string {
	if msg, ok := ErrorInfo[code]; ok {
		return msg
	}
	return "未知错误"
}
