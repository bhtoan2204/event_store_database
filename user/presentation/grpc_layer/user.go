package grpc_layer

import (
	"context"
	"event_sourcing_user/proto/user"
)

func (g *grpcPresentation) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	return nil, nil
}
