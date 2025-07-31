package grpc_layer

import (
	"context"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/proto/payment"
)

type ITransactionGrpcService interface {
	CreateTransaction(ctx context.Context, req *payment.CreateTransactionRequest) (*payment.CreateTransactionResponse, error)
}

type TransactionGrpcService struct {
	transactionUsecase usecase.ITransactionUsecase
}

func NewTransactionGrpcService(transactionUsecase usecase.ITransactionUsecase) *TransactionGrpcService {
	return &TransactionGrpcService{transactionUsecase: transactionUsecase}
}

func (s *TransactionGrpcService) CreateTransaction(ctx context.Context, req *payment.CreateTransactionRequest) (*payment.CreateTransactionResponse, error) {
	return nil, nil
}
