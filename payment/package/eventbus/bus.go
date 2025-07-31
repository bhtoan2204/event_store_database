package eventbus

import (
	"context"
	"sync"
)

type EventBus interface {
	RegisterHandler(event IEvent, handler EventHandler) error
	Dispatch(ctx context.Context, event IEvent) error
}

type eventBus struct {
	handlers map[string]EventHandler
	mu       sync.RWMutex
}

func NewEventBus() EventBus {
	return &eventBus{
		handlers: make(map[string]EventHandler),
	}
}

func (cb *eventBus) RegisterHandler(event IEvent, handler EventHandler) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if _, exists := cb.handlers[event.EventName()]; exists {
		return ErrHandlerAlreadyRegistered
	}

	cb.handlers[event.EventName()] = handler
	return nil
}

func (cb *eventBus) Dispatch(ctx context.Context, event IEvent) error {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	handler, exists := cb.handlers[event.EventName()]
	if !exists {
		return ErrHandlerNotFound
	}

	return handler.Handle(ctx, event)
}
