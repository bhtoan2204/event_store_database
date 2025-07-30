package commandbus

import "context"

type CommandHandler interface {
	Handle(ctx context.Context, command ICommand) error
}
