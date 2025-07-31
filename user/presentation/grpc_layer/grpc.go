package grpc_layer

import (
	"event_sourcing_user/constant"
	"event_sourcing_user/domain/usecase"
	"event_sourcing_user/infrastructure/persistent/repository"

	"google.golang.org/grpc"
)

type GrpcPresentation interface {
	AuthGrpcService
	UserGrpcService
	GRPCServer() *grpc.Server
}
type grpcPresentation struct {
	server *grpc.Server

	AuthGrpcService
	UserGrpcService
}

func NewGrpcPresentation(s *grpc.Server, rf repository.IRepositoryFactory, config *constant.Config, usecase usecase.Usecase) (GrpcPresentation, error) {
	return &grpcPresentation{
		server:          s,
		AuthGrpcService: NewAuthGrpcService(usecase.AuthUsecase()),
		UserGrpcService: NewUserGrpcService(usecase.UserUsecase()),
	}, nil
}

func (g *grpcPresentation) GRPCServer() *grpc.Server { return g.server }
