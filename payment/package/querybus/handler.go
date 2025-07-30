package querybus

import "context"

type QueryHandler interface {
	Handle(ctx context.Context, query IQuery) error
}
