package grpc_layer

import (
	"context"
	"event_sourcing_user/application/command"
	"event_sourcing_user/domain/usecase"
	"event_sourcing_user/package/ierror"
	"event_sourcing_user/package/logger"
	"event_sourcing_user/proto/user"
	"fmt"

	"go.uber.org/zap"
)

type AuthService interface {
	Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error)
	Refresh(ctx context.Context, req *user.RefreshTokenRequest) (*user.RefreshTokenResponse, error)
}

type authService struct {
	usecase usecase.AuthUsecase
}

func NewAuthService(usecase usecase.AuthUsecase) AuthService {
	return &authService{usecase: usecase}
}

func (g *authService) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	log := logger.FromContext(ctx)
	command := command.NewLoginCommand(req)
	if err := command.Validate(); err != nil {
		log.Error("Invalid request", zap.Error(err))
		return nil, ierror.Error(err)
	}
	output, err := g.usecase.Login(ctx, command)
	if err != nil {
		log.Error("Error logging in", zap.Error(err))
		return nil, ierror.Error(err)
	}
	return &user.LoginResponse{
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	}, nil
}

func (g *authService) Refresh(ctx context.Context, req *user.RefreshTokenRequest) (*user.RefreshTokenResponse, error) {
	fmt.Println("Refresh request received")
	return nil, nil
}
