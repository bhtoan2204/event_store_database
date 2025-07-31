package event_listener_layer

import "context"

type EventListener interface {
	Start(ctx context.Context) error
}
