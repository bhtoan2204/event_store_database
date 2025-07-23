package grpc_layer

import (
	"context"
	"event_sourcing_user/proto/user"
	"fmt"
)

func (g *grpcPresentation) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	fmt.Println("CreateUser request received")
	return nil, nil
}
