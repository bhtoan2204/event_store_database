package command

import (
	"context"
	"event_sourcing_payment/application/command/transaction_command"
	"event_sourcing_payment/application/command/user_command"
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
	cb.RegisterHandler(user_command.CreateUserCommand{}, user_command.NewCreateUserHandler(useCase))

	cb.RegisterHandler(transaction_command.CreateTransactionCommand{}, transaction_command.NewCreateTransactionHandler(useCase.TransactionUsecase()))

	return &CommandBus{
		CommandBus: cb,
		esdbStorer: esdbStorer,
		useCase:    useCase,
	}
}

func (c *CommandBus) Dispatch(ctx context.Context, command commandbus.ICommand) error {
	return c.CommandBus.Dispatch(ctx, command)
}
