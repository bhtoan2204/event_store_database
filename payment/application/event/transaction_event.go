package event

import (
	"context"
	"event_sourcing_payment/package/eventbus"
	"fmt"
)

var _ eventbus.IEvent = (*TransactionCreatedEvent)(nil)

type TransactionCreatedEvent struct {
	AccountID string
	Amount    float64
	Type      string // deposit, withdraw, transfer
	Reference string
}

func (e TransactionCreatedEvent) EventName() string {
	return "TransactionCreated"
}

type DomainEvent interface {
	EventName() string
}

type TransactionCreatedHandler struct{}

func NewTransactionCreatedHandler() eventbus.EventHandler {
	return &TransactionCreatedHandler{}
}

func (h *TransactionCreatedHandler) Handle(ctx context.Context, e eventbus.IEvent) error {
	evt, ok := e.(TransactionCreatedEvent)
	if !ok {
		return fmt.Errorf("invalid event type")
	}

	// TODO: implement processing logic, example:
	fmt.Printf("Processing transaction event: %+v\n", evt)
	return nil
}
