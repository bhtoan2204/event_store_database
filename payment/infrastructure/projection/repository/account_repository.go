package repository

import (
	"context"
	"event_sourcing_payment/infrastructure/projection/persistent_object"

	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccountByID(ctx context.Context, id string) (*persistent_object.Account, error)
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(ctx context.Context, db *gorm.DB) IAccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, id string) (*persistent_object.Account, error) {
	return nil, nil
}
