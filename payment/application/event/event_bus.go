package event

import (
	"event_sourcing_payment/package/eventbus"
)

type EventBus struct {
	eventbus.EventBus
}

func NewEventBus() *EventBus {
	eb := eventbus.NewEventBus()

	eb.RegisterHandler(TransactionCreatedEvent{}, NewTransactionCreatedHandler())
	return &EventBus{
		EventBus: eb,
	}
}
