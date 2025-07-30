package commandbus

import (
	"context"
	"sync"
)

type CommandBus interface {
	RegisterHandler(command ICommand, handler CommandHandler) error
	Dispatch(ctx context.Context, command ICommand) error
}

type commandBus struct {
	handlers map[string]CommandHandler
	mu       sync.RWMutex
}

func NewCommandBus() CommandBus {
	return &commandBus{
		handlers: make(map[string]CommandHandler),
	}
}

func (cb *commandBus) RegisterHandler(command ICommand, handler CommandHandler) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if _, exists := cb.handlers[command.CommandName()]; exists {
		return ErrHandlerAlreadyRegistered
	}

	cb.handlers[command.CommandName()] = handler
	return nil
}

func (cb *commandBus) Dispatch(ctx context.Context, command ICommand) error {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	handler, exists := cb.handlers[command.CommandName()]
	if !exists {
		return ErrHandlerNotFound
	}

	return handler.Handle(ctx, command)
}
