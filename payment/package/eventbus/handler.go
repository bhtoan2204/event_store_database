package eventbus

import "context"

type EventHandler interface {
	Handle(ctx context.Context, event IEvent) error
}
