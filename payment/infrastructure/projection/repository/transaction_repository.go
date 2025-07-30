package repository

import (
	"context"
	"event_sourcing_payment/infrastructure/projection/persistent_object"
)

type ITransactionRepository interface {
	GetTransactionByAccountID(ctx context.Context, accountID string) (*[]persistent_object.Transaction, error)
}

type TransactionRepository struct {
}

func NewTransactionRepository() ITransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) GetTransactionByAccountID(ctx context.Context, accountID string) (*[]persistent_object.Transaction, error) {
	return nil, nil
}
