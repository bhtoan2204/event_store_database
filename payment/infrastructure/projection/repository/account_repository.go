package repository

import (
	"context"
	"event_sourcing_payment/infrastructure/projection/persistent_object"
)

type IAccountRepository interface {
	GetAccountByID(ctx context.Context, id string) (*persistent_object.Account, error)
}

type AccountRepository struct {
}

func NewAccountRepository() IAccountRepository {
	return &AccountRepository{}
}

func (r *AccountRepository) GetAccountByID(ctx context.Context, id string) (*persistent_object.Account, error) {
	return nil, nil
}
