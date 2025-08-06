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
	var transactions []persistent_object.Transaction
	if err := r.db.WithContext(ctx).Where("account_no = ?", accountNo).Order("created_at desc").Find(&transactions).Error; err != nil {
		return nil, err
	}
	return &transactions, nil
}
