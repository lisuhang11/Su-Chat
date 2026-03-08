package userRpcClient

import (
	"Su-chat/services/user/gen"
	"google.golang.org/grpc"
)

import (
	"context"
)

// 类型别名：简化其他模块的调用（无需写 gen.XXX）
type (
	UserRegisterReq       = gen.UserRegisterReq
	UserRegisterResp      = gen.UserRegisterResp
	UserLoginReq          = gen.UserLoginReq
	UserLoginResp         = gen.UserLoginResp
	UserInfoReq           = gen.UserInfoReq
	UserInfoResp          = gen.UserInfoResp
	UserInfoUpdateReq     = gen.UserInfoUpdateReq
	UserInfoUpdateResp    = gen.UserInfoUpdateResp
	UsersOnlineStatusReq  = gen.UsersOnlineStatusReq
	UsersOnlineStatusResp = gen.UsersOnlineStatusResp
	UserOnlineStatus      = gen.UserOnlineStatus
)

// User 定义对外暴露的客户端接口，包含所有 UserService 的 RPC 方法
// 接口风格和原示例一致，但基于原生 gRPC 实现
type User interface {
	// 用户注册
	Register(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	// 用户登录
	Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error)
	// 查询用户信息
	GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	// 更新用户信息
	UpdateUserInfo(ctx context.Context, in *UserInfoUpdateReq, opts ...grpc.CallOption) (*UserInfoUpdateResp, error)
	// 批量查询用户在线状态
	GetUsersOnlineStatus(ctx context.Context, in *UsersOnlineStatusReq, opts ...grpc.CallOption) (*UsersOnlineStatusResp, error)
}

// defaultUser 是 User 接口的默认实现（纯 gRPC 版本）
type defaultUser struct {
	conn grpc.ClientConnInterface // 原生 gRPC 连接接口
}

// NewUser 创建 User 客户端实例（供其他模块调用）
// 参数为原生 gRPC 连接，支持任意符合 grpc.ClientConnInterface 的连接（如 *grpc.ClientConn）
func NewUser(conn grpc.ClientConnInterface) User {
	return &defaultUser{
		conn: conn,
	}
}

// Register 封装用户注册 RPC 方法
func (m *defaultUser) Register(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	// 创建 gen 包的 UserService 客户端（原生 gRPC）
	client := gen.NewUserServiceClient(m.conn)
	// 调用底层 RPC 方法
	return client.Register(ctx, in, opts...)
}

// Login 封装用户登录 RPC 方法
func (m *defaultUser) Login(ctx context.Context, in *UserLoginReq, opts ...grpc.CallOption) (*UserLoginResp, error) {
	client := gen.NewUserServiceClient(m.conn)
	return client.Login(ctx, in, opts...)
}

// GetUserInfo 封装查询用户信息 RPC 方法
func (m *defaultUser) GetUserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	client := gen.NewUserServiceClient(m.conn)
	return client.GetUserInfo(ctx, in, opts...)
}

// UpdateUserInfo 封装更新用户信息 RPC 方法
func (m *defaultUser) UpdateUserInfo(ctx context.Context, in *UserInfoUpdateReq, opts ...grpc.CallOption) (*UserInfoUpdateResp, error) {
	client := gen.NewUserServiceClient(m.conn)
	return client.UpdateUserInfo(ctx, in, opts...)
}

// GetUsersOnlineStatus 封装批量查询在线状态 RPC 方法
func (m *defaultUser) GetUsersOnlineStatus(ctx context.Context, in *UsersOnlineStatusReq, opts ...grpc.CallOption) (*UsersOnlineStatusResp, error) {
	client := gen.NewUserServiceClient(m.conn)
	return client.GetUsersOnlineStatus(ctx, in, opts...)
}
