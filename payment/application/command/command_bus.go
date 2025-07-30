package command

import (
	"event_sourcing_payment/application/command/user"
	"event_sourcing_payment/package/commandbus"
)

type CommandBus struct {
	commandbus.CommandBus
}

func NewCommandBus() *CommandBus {
	cb := commandbus.NewCommandBus()
	cb.RegisterHandler(user.CreateUserCommand{}, &user.CreateUserHandler{})
	return &CommandBus{
		CommandBus: cb,
	}
}
