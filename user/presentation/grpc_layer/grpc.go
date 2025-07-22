package grpc_layer

import (
	"event_sourcing_user/infrastructure/persistent/repository"
	"event_sourcing_user/proto/user"

	"google.golang.org/grpc"
)

type GrpcPresentation interface {
	user.UserServiceServer
}
type grpcPresentation struct {
	server *grpc.Server
}

func NewGrpcPresentation(server *grpc.Server, repositoryFactory repository.IRepositoryFactory) GrpcPresentation {
	return &grpcPresentation{server: server}
}
