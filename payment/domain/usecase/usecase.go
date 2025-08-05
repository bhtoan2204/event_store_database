package usecase

import "event_sourcing_payment/infrastructure/eventstore/esdb_storer"

type IUseCase interface {
	AccountUsecase() IAccountUsecase
	TransactionUsecase() ITransactionUsecase
}

type UseCase struct {
	accountUsecase     IAccountUsecase
	transactionUsecase ITransactionUsecase
}

func NewUseCase(esdbStorer esdb_storer.IEventStorer) IUseCase {
	return &UseCase{
		accountUsecase:     NewAccountUsecase(),
		transactionUsecase: NewTransactionUsecase(esdbStorer),
	}
}

func (u *UseCase) AccountUsecase() IAccountUsecase {
	return u.accountUsecase
}

func (u *UseCase) TransactionUsecase() ITransactionUsecase {
	return u.transactionUsecase
}
