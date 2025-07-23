package grpc_layer

import (
	"event_sourcing_user/constant"
	"event_sourcing_user/domain/usecase"
	"event_sourcing_user/infrastructure/persistent/repository"

	"google.golang.org/grpc"
)

type GrpcPresentation interface {
	AuthService
	UserService
	GRPCServer() *grpc.Server
}
type grpcPresentation struct {
	server *grpc.Server

	AuthService
	UserService
}

func NewGrpcPresentation(s *grpc.Server, rf repository.IRepositoryFactory, config *constant.Config, usecase usecase.Usecase) (GrpcPresentation, error) {
	return &grpcPresentation{
		server:      s,
		AuthService: NewAuthService(usecase.AuthUsecase()),
		UserService: NewUserService(usecase.UserUsecase()),
	}, nil
}

func (g *grpcPresentation) GRPCServer() *grpc.Server { return g.server }
