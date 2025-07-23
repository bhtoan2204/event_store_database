package grpc_layer

import (
	"context"
	"event_sourcing_user/application/command"
	"event_sourcing_user/domain/usecase"
	"event_sourcing_user/package/ierror"
	"event_sourcing_user/package/logger"
	"event_sourcing_user/proto/user"

	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error)
}

type userService struct {
	usecase usecase.UserUsecase
}

func NewUserService(usecase usecase.UserUsecase) UserService {
	return &userService{usecase: usecase}
}

func (g *userService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	log := logger.FromContext(ctx)
	command := command.NewCreateUser(req)
	if err := command.Validate(); err != nil {
		log.Error("Invalid request", zap.Error(err))
		return nil, ierror.Error(err)
	}
	err := g.usecase.CreateUser(ctx, command)
	if err != nil {
		log.Error("Error creating user", zap.Error(err))
		return nil, ierror.Error(err)
	}

	return &user.CreateUserResponse{
		Success: true,
	}, nil
}
