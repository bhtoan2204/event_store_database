package event

import (
	"event_sourcing_payment/application/event/transaction"
	event "event_sourcing_payment/application/event/transaction"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/package/eventbus"
)

type EventBus struct {
	eventbus.EventBus
	useCase usecase.IUseCase
}

func NewEventBus(useCase usecase.IUseCase) *EventBus {
	eb := eventbus.NewEventBus()
	eb.RegisterHandler(event.TransactionCreatedEvent{}, transaction.NewTransactionCreatedHandler(useCase))

	return &EventBus{
		EventBus: eb,
		useCase:  useCase,
	}
}
