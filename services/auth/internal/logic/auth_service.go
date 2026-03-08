package logic

import (
	"Su-chat/pkg/tools"
	"Su-chat/services/auth/gen"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServer struct {
	gen.UnimplementedAuthServiceServer
}

// NewAuthServer 创建 authServer 实例
func NewAuthServer() *authServer {
	return &authServer{}
}

// 为指定用户生成新的访问令牌和刷新令牌
func (c *authServer) IssueToken(ctx context.Context, in *gen.IssueTokenRequest) (*gen.IssueTokenResponse, error) {
	var userId = in.GetUserId()
	if userId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "refresh token is required")
	}
	var access_token, _ = tools.GenerateToken("Su-chat-secret-key", userId, 7200)
	var refresh_token, _ = tools.GenerateToken("Su-chat-secret-key", userId, 86400)
	var rsp = &gen.IssueTokenResponse{
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ExpiresIn:    7200,
	}
	return rsp, nil
}

// 用有效的刷新令牌换取新的访问令牌
func (c *authServer) RefreshToken(ctx context.Context, in *gen.RefreshTokenRequest) (*gen.RefreshTokenResponse, error) {
	var refresh_token = in.GetRefreshToken()
	if refresh_token == "" {
		return nil, status.Errorf(codes.InvalidArgument, "refresh token is required")
	}
	var claim, err = tools.ParseToken(refresh_token, "Su-chat-secret-key")
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "refresh token is invalid")
	}
	// 生成新的访问令牌
	newAccessToken, err := tools.GenerateToken("Su-chat-secret-key", claim.UserID, 7200)
	// 生成新的刷新令牌
	newRefreshToken, err := tools.GenerateToken("Su-chat-secret-key", claim.UserID, 86400)
	return &gen.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    7200,
	}, nil
}
