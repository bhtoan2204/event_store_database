package command

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, command ICommand) error
}

type commandHandler struct {
	command ICommand
}

func NewCommandHandler(command ICommand) CommandHandler {
	return &commandHandler{
		command: command,
	}
}

func (h *commandHandler) Handle(ctx context.Context, command ICommand) error {
	return nil
}
