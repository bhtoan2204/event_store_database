package usecase

import (
	"context"
	"errors"
	"event_sourcing_payment/constant"
	"event_sourcing_payment/domain/aggregation"
	"event_sourcing_payment/dto"
	"event_sourcing_payment/infrastructure/eventstore/esdb_storer"
	"event_sourcing_payment/infrastructure/projection/repository"
)

type ITransactionUsecase interface {
	CreateTransaction(ctx context.Context, req *dto.CreateTransactionRequestDto) error
	ListTransaction(ctx context.Context, req *dto.ListTransactionRequestDto) (*dto.ListTransactionResponseDto, error)
}

type TransactionUsecase struct {
	eventStorer esdb_storer.IEventStorer
	repoFactory repository.IFactoryRepository
}

func NewTransactionUsecase(storer esdb_storer.IEventStorer, repoFactory repository.IFactoryRepository) ITransactionUsecase {
	return &TransactionUsecase{
		eventStorer: storer,
		repoFactory: repoFactory,
	}
}

func (u *TransactionUsecase) CreateTransaction(ctx context.Context, req *dto.CreateTransactionRequestDto) error {
	if err := req.Validate(); err != nil {
		return err
	}

	agg := aggregation.NewAccountAggregate(req.AccountNo)

	switch req.Type {
	case constant.TransactionTypeDeposit.String():
		agg.Deposit(req.Amount, req.Reference)
	case constant.TransactionTypeWithdraw.String():
		agg.Withdraw(req.Amount, req.Reference)
	default:
		return errors.New("unsupported transaction type")
	}

	return u.eventStorer.Append(ctx, agg.ID, agg.GetUncommittedEvents())
}

func (u *TransactionUsecase) ListTransaction(ctx context.Context, req *dto.ListTransactionRequestDto) (*dto.ListTransactionResponseDto, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Get transactions from projection database
	transactions, err := u.repoFactory.TransactionRepository().GetTransactionByAccountNo(ctx, req.AccountNo)
	if err != nil {
		return nil, err
	}

	if transactions == nil {
		return &dto.ListTransactionResponseDto{
			Rows:       []*aggregation.TransactionAggregate{},
			TotalCount: 0,
			TotalPages: 0,
			Page:       0,
			PageSize:   0,
		}, nil
	}

	// Convert persistent objects to DTOs
	var responseDtos []*aggregation.TransactionAggregate
	for _, transaction := range *transactions {
		responseDtos = append(responseDtos, &aggregation.TransactionAggregate{
			ID:              transaction.ID,
			TransactionCode: transaction.TransactionCode,
			AccountNo:       transaction.AccountNo,
			Type:            transaction.Type.String(),
			Amount:          float64(transaction.Amount),
			Reference:       transaction.Reference,
		})
	}

	return &dto.ListTransactionResponseDto{
		Rows:       responseDtos,
		TotalCount: len(responseDtos),
		TotalPages: 1,
		Page:       1,
		PageSize:   1,
	}, nil
}
