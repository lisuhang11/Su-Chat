package models

// 用户注册请求
type UserRegisterReq struct {
	NickName     string            `json:"nickname"`      // 必填，昵称
	Password     string            `json:"password"`      // 必填，密码
	Phone        string            `json:"phone"`         // 选填，手机号（唯一）
	Email        string            `json:"email"`         // 选填，邮箱（唯一）
	UserPortrait string            `json:"user_portrait"` // 选填，头像 URL
	ExtFields    map[string]string `json:"ext_fields"`    // 选填，扩展字段（存入 userexts 表）
}

// 用户注册响应
type UserRegisterResp struct {
	UserId        string `json:"user_id"`       // 用户ID
	Access_Token  string `json:"access_token"`  // 访问 Token
	Refresh_Token string `json:"refresh_token"` // 刷新 Token
}

// 用户登录请求
type UserLoginReq struct {
	UserId   string `json:"user_id"`  // 必填，登录标识（支持账号/手机号/邮箱）
	Password string `json:"password"` // 必填，密码
}

// 用户登录响应
type UserLoginResp struct {
	UserId        string `json:"user_id"`       // 用户内部 ID
	Access_Token  string `json:"access_token"`  // 访问 Token
	Refresh_Token string `json:"refresh_token"` // 刷新 Token
}

// 查询用户信息请求
type UserInfoReq struct {
	UserId string `json:"user_id"` // 必填，用户内部 ID
}

// 查询用户信息响应
type UserInfoResp struct {
	UserId       string            `json:"user_id"`       // 用户内部 ID
	NickName     string            `json:"nickname"`      // 用户昵称
	Phone        string            `json:"phone"`         // 手机号
	Email        string            `json:"email"`         // 邮箱
	UserPortrait string            `json:"user_portrait"` // 头像 URL
	ExtFields    map[string]string `json:"ext_fields"`    // 扩展字段
}

// 更新用户信息请求
type UserInfoUpdateReq struct {
	UserId       string            `json:"user_id"`       // 必填，用户内部 ID
	NickName     string            `json:"nickname"`      // 选填，新昵称
	Phone        string            `json:"phone"`         // 选填，新手机号
	Email        string            `json:"email"`         // 选填，新邮箱
	UserPortrait string            `json:"user_portrait"` // 选填，新头像
	ExtFields    map[string]string `json:"ext_fields"`    // 选填，更新扩展字段（全量替换）
}

// 更新用户信息响应
type UserInfoUpdateResp struct {
	UserId       string            `json:"user_id"`
	Account      string            `json:"account"`
	NickName     string            `json:"nickname"`
	Phone        string            `json:"phone"`
	Email        string            `json:"email"`
	UserPortrait string            `json:"user_portrait"`
	ExtFields    map[string]string `json:"ext_fields"`
}

// 用户在线状态查询请求
type UsersOnlineStatusReq struct {
	UserIDs []string `json:"user_ids"` // 需要查询的用户内部 ID 列表
}

// 单个用户在线状态
type UserOnlineStatus struct {
	UserId   string `json:"user_id"`   // 用户内部 ID
	IsOnline bool   `json:"is_online"` // 是否在线
}

// 用户在线状态响应
type UserOnlineStatusResp struct {
	Items []UserOnlineStatus `json:"items"`
}
