package grpc_layer

import (
	"event_sourcing_payment/constant"

	"google.golang.org/grpc"
)

type GrpcPresentation interface {
	GRPCServer() *grpc.Server
}
type grpcPresentation struct {
	server *grpc.Server
}

func NewGrpcPresentation(s *grpc.Server, config *constant.Config) (GrpcPresentation, error) {
	return &grpcPresentation{
		server: s,
	}, nil
}

func (g *grpcPresentation) GRPCServer() *grpc.Server { return g.server }
