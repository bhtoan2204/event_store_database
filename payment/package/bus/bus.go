package command

type CommandBus struct {
	handlers map[string]CommandHandler
}
