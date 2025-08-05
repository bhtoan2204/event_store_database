package grpc_layer

import (
	"event_sourcing_payment/application/command"

	"google.golang.org/grpc"
)

type GrpcPresentation interface {
	ITransactionGrpcService
	GRPCServer() *grpc.Server
}
type grpcPresentation struct {
	server *grpc.Server
	ITransactionGrpcService
}

func NewGrpcPresentation(
	s *grpc.Server,
	commandBus *command.CommandBus,
) (GrpcPresentation, error) {
	transactionService := NewTransactionGrpcService(commandBus)
	return &grpcPresentation{
		server:                  s,
		ITransactionGrpcService: transactionService,
	}, nil
}

func (g *grpcPresentation) GRPCServer() *grpc.Server { return g.server }
