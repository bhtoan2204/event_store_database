package broker

import (
	"context"
	"event_sourcing_payment/constant"
)

type Producer interface {
	Produce(ctx context.Context, topic string, message []byte) error
	Close(ctx context.Context)
}

func NewProducer(config *constant.Config) (Producer, error) {
	return nil, nil
}
