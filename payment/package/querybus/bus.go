package querybus

import (
	"context"
	"sync"
)

type QueryBus interface {
	RegisterHandler(query IQuery, handler QueryHandler) error
	Dispatch(ctx context.Context, query IQuery) error
}

type queryBus struct {
	handlers map[string]QueryHandler
	mu       sync.RWMutex
}

func NewQueryBus() QueryBus {
	return &queryBus{
		handlers: make(map[string]QueryHandler),
	}
}

func (cb *queryBus) RegisterHandler(query IQuery, handler QueryHandler) error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if _, exists := cb.handlers[query.QueryName()]; exists {
		return ErrHandlerAlreadyRegistered
	}

	cb.handlers[query.QueryName()] = handler
	return nil
}

func (cb *queryBus) Dispatch(ctx context.Context, query IQuery) error {
	cb.mu.RLock()
	defer cb.mu.RUnlock()

	handler, exists := cb.handlers[query.QueryName()]
	if !exists {
		return ErrHandlerNotFound
	}

	return handler.Handle(ctx, query)
}
