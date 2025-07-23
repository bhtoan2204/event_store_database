package grpc_layer

import (
	"context"
	"event_sourcing_user/proto/user"
	"fmt"
)

func (g *grpcPresentation) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	fmt.Println("Login request received")
	return nil, nil
}

func (g *grpcPresentation) Refresh(ctx context.Context, req *user.RefreshTokenRequest) (*user.RefreshTokenResponse, error) {
	fmt.Println("Refresh request received")
	return nil, nil
}
