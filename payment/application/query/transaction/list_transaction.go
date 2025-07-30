package transaction

import (
	"context"
	"event_sourcing_payment/package/querybus"
	"fmt"
)

type ListTransactionQuery struct {
	AccountID string
	Page      int
	PageSize  int
}

func (q ListTransactionQuery) QueryName() string {
	return "ListTransactionQuery"
}

type ListTransactionHandler struct {
}

func (h *ListTransactionHandler) Handle(ctx context.Context, query querybus.IQuery) error {
	q, ok := query.(ListTransactionQuery)
	if !ok {
		return fmt.Errorf("invalid query type")
	}
	fmt.Println("List transaction:", q.AccountID, q.Page, q.PageSize)
	return nil
}
