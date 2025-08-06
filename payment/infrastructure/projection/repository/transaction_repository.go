package repository

import (
	"context"
	"event_sourcing_payment/infrastructure/projection/persistent_object"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	GetTransactionByAccountNo(ctx context.Context, accountNo string) (*[]persistent_object.Transaction, error)
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(ctx context.Context, db *gorm.DB) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) GetTransactionByAccountNo(ctx context.Context, accountNo string) (*[]persistent_object.Transaction, error) {
	return nil, nil
}
