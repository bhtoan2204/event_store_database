package usecase

import (
	"context"
	"errors"
	"event_sourcing_payment/domain/aggregation"
	"event_sourcing_payment/dto"
	"event_sourcing_payment/infrastructure/eventstore/esdb_storer"
)

type ITransactionUsecase interface {
	CreateTransaction(ctx context.Context, req *dto.CreateTransactionRequestDto) error
}

type TransactionUsecase struct {
	eventStorer esdb_storer.IEventStorer
}

func NewTransactionUsecase(storer esdb_storer.IEventStorer) ITransactionUsecase {
	return &TransactionUsecase{eventStorer: storer}
}

func (u *TransactionUsecase) CreateTransaction(ctx context.Context, req *dto.CreateTransactionRequestDto) error {
	if err := req.Validate(); err != nil {
		return err
	}

	agg := aggregation.NewAccountAggregate(req.AccountNo)

	switch req.Type {
	case "deposit":
		agg.Deposit(req.Amount, req.Reference)
	case "withdraw":
		agg.Withdraw(req.Amount, req.Reference)
	default:
		return errors.New("unsupported transaction type")
	}

	return u.eventStorer.Append(ctx, agg.ID, agg.GetUncommittedEvents())
}
