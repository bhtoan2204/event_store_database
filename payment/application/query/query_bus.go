package query

import (
	"event_sourcing_payment/application/query/transaction"
	"event_sourcing_payment/package/querybus"
)

type QueryBus struct {
	querybus.QueryBus
}

func NewQueryBus() *QueryBus {
	cb := querybus.NewQueryBus()
	cb.RegisterHandler(transaction.ListTransactionQuery{}, &transaction.ListTransactionHandler{})
	return &QueryBus{
		QueryBus: cb,
	}
}
