package grpc_layer

import (
	"context"
	"event_sourcing_payment/application/command"
	"event_sourcing_payment/application/command/transaction_command"
	"event_sourcing_payment/proto/payment"
)

type ITransactionGrpcService interface {
	CreateTransaction(ctx context.Context, req *payment.CreateTransactionRequest) (*payment.CreateTransactionResponse, error)
}

type TransactionGrpcService struct {
	commandBus *command.CommandBus
}

func NewTransactionGrpcService(commandBus *command.CommandBus) *TransactionGrpcService {
	return &TransactionGrpcService{commandBus: commandBus}
}

func (s *TransactionGrpcService) CreateTransaction(ctx context.Context, req *payment.CreateTransactionRequest) (*payment.CreateTransactionResponse, error) {
	createTransactionCommand := transaction_command.CreateTransactionCommand{
		AccountNo: req.AccountNo,
		Amount:    req.Amount,
		Type:      req.Type,
		Reference: req.Reference,
	}

	if err := s.commandBus.Dispatch(ctx, createTransactionCommand); err != nil {
		return nil, err
	}

	return &payment.CreateTransactionResponse{
		TransactionCode: "123",
	}, nil
}
