package command

import (
	command "event_sourcing_payment/application/command/user"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/infrastructure/eventstore/esdb_storer"
	"event_sourcing_payment/package/commandbus"
)

type CommandBus struct {
	commandbus.CommandBus
	esdbStorer esdb_storer.IEventStorer
	useCase    usecase.IUseCase
}

func NewCommandBus(esdbStorer esdb_storer.IEventStorer, useCase usecase.IUseCase) *CommandBus {
	cb := commandbus.NewCommandBus()
	cb.RegisterHandler(command.CreateUserCommand{}, command.NewCreateUserHandler(useCase))
	return &CommandBus{
		CommandBus: cb,
		esdbStorer: esdbStorer,
		useCase:    useCase,
	}
}
