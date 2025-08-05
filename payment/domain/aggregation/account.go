package aggregation

import (
	"event_sourcing_payment/application/event"
	"event_sourcing_payment/constant"
)

type AccountAggregate struct {
	ID     string
	Events []event.DomainEvent
}

func NewAccountAggregate(id string) *AccountAggregate {
	return &AccountAggregate{ID: id}
}

func (a *AccountAggregate) Deposit(amount float64, ref string) {
	e := event.TransactionCreatedEvent{
		AccountID: a.ID,
		Amount:    amount,
		Type:      constant.TransactionTypeDeposit.String(),
		Reference: ref,
	}
	a.Events = append(a.Events, e)
}

func (a *AccountAggregate) Withdraw(amount float64, ref string) {
	e := event.TransactionCreatedEvent{
		AccountID: a.ID,
		Amount:    amount,
		Type:      constant.TransactionTypeWithdraw.String(),
		Reference: ref,
	}
	a.Events = append(a.Events, e)
}

func (a *AccountAggregate) GetUncommittedEvents() []event.DomainEvent {
	return a.Events
}
