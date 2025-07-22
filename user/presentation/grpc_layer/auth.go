package grpc_layer

import (
	"context"
	"event_sourcing_user/proto/user"
)

func (g *grpcPresentation) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	return nil, nil
}

func (g *grpcPresentation) Refresh(ctx context.Context, req *user.RefreshTokenRequest) (*user.RefreshTokenResponse, error) {
	return nil, nil
}
