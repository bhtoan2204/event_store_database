package transaction

import (
	"context"
	"event_sourcing_payment/constant"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/package/eventbus"
	"fmt"
)

type TransactionCreatedEvent struct {
	AccountID uint
	Amount    int64
	Type      constant.TransactionType
	Reference string
}

func (e TransactionCreatedEvent) EventName() string {
	return "TransactionCreated"
}

func NewTransactionCreatedEvent(accountID uint, amount int64, typ constant.TransactionType, ref string) *TransactionCreatedEvent {
	return &TransactionCreatedEvent{
		AccountID: accountID,
		Amount:    amount,
		Type:      typ,
		Reference: ref,
	}
}

type TransactionCreatedHandler struct {
	useCase usecase.IUseCase
}

func NewTransactionCreatedHandler(useCase usecase.IUseCase) *TransactionCreatedHandler {
	return &TransactionCreatedHandler{useCase: useCase}
}

func (h *TransactionCreatedHandler) Handle(ctx context.Context, event eventbus.IEvent) error {
	e, ok := event.(*TransactionCreatedEvent)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	fmt.Println("Transaction created", e.AccountID, e.Amount, e.Type, e.Reference)

	return nil
}

var _ eventbus.IEvent = (*TransactionCreatedEvent)(nil)
