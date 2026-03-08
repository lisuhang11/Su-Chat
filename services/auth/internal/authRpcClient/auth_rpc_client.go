package auth_rpc_client

import (
	"Su-chat/services/auth/gen"
	"context"
	"google.golang.org/grpc"
)

// 类型别名（可选，方便调用方）
type (
	IssueTokenReq    = gen.IssueTokenRequest
	IssueTokenResp   = gen.IssueTokenResponse
	RefreshTokenReq  = gen.RefreshTokenRequest
	RefreshTokenResp = gen.RefreshTokenResponse
)

type Auth interface {
	IssueToken(ctx context.Context, in *IssueTokenReq, opts ...grpc.CallOption) (*IssueTokenResp, error)
	RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error)
}

type defaultAuth struct {
	conn grpc.ClientConnInterface
}

func NewAuth(conn grpc.ClientConnInterface) Auth {
	return &defaultAuth{conn: conn}
}

func (d *defaultAuth) IssueToken(ctx context.Context, in *IssueTokenReq, opts ...grpc.CallOption) (*IssueTokenResp, error) {
	client := gen.NewAuthServiceClient(d.conn)
	return client.IssueToken(ctx, in, opts...)
}

func (d *defaultAuth) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*RefreshTokenResp, error) {
	client := gen.NewAuthServiceClient(d.conn)
	return client.RefreshToken(ctx, in, opts...)
}
