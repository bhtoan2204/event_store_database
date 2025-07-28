package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

var _ Consumer = (*consumer)(nil)

type Consumer interface {
}

type consumer struct {
	ins *kafka.Consumer
}

func NewConsumer() Consumer {
	return &consumer{}
}
