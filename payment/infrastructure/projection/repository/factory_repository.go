package repository

import (
	"context"
	"event_sourcing_payment/infrastructure/projection"
)

type IFactoryRepository interface {
	WithTransaction(ctx context.Context, fn func(IFactoryRepository) error) error
	AccountRepository() IAccountRepository
	TransactionRepository() ITransactionRepository
}

type FactoryRepository struct {
	projectionConnection  *projection.ProjectionConnection
	accountRepository     IAccountRepository
	transactionRepository ITransactionRepository
}

func NewFactoryRepository(ctx context.Context, projectionConnection *projection.ProjectionConnection) IFactoryRepository {
	accountRepository := NewAccountRepository(ctx, projectionConnection.GetDB())
	transactionRepository := NewTransactionRepository(ctx, projectionConnection.GetDB())
	return &FactoryRepository{
		projectionConnection:  projectionConnection,
		accountRepository:     accountRepository,
		transactionRepository: transactionRepository,
	}
}

func (r *FactoryRepository) WithTransaction(ctx context.Context, fn func(IFactoryRepository) error) (err error) {
	tx := r.projectionConnection.GetDB().Begin()
	tr := NewFactoryRepository(ctx, r.projectionConnection)

	err = tx.Error
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = fn(tr)

	return err
}

func (f *FactoryRepository) AccountRepository() IAccountRepository {
	return f.accountRepository
}

func (f *FactoryRepository) TransactionRepository() ITransactionRepository {
	return f.transactionRepository
}
