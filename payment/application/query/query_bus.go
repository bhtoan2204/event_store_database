package query

import (
	query "event_sourcing_payment/application/query/transaction"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/infrastructure/projection/repository"
	"event_sourcing_payment/package/querybus"
)

type QueryBus struct {
	querybus.QueryBus
	factoryRepository repository.IFactoryRepository
	useCase           usecase.IUseCase
}

func NewQueryBus(factoryRepository repository.IFactoryRepository, useCase usecase.IUseCase) *QueryBus {
	cb := querybus.NewQueryBus()
	cb.RegisterHandler(query.ListTransactionQuery{}, query.NewListTransactionHandler(useCase))
	return &QueryBus{
		QueryBus:          cb,
		factoryRepository: factoryRepository,
		useCase:           useCase,
	}
}
