package grpc_layer

import (
	"event_sourcing_payment/constant"
	"event_sourcing_payment/domain/usecase"

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
	config *constant.Config,
	usecase usecase.IUseCase,
) (GrpcPresentation, error) {
	transactionService := NewTransactionGrpcService(usecase.TransactionUsecase())
	return &grpcPresentation{
		server:                  s,
		ITransactionGrpcService: transactionService,
	}, nil
}

func (g *grpcPresentation) GRPCServer() *grpc.Server { return g.server }
