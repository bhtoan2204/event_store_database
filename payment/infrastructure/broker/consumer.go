package broker

import (
	"event_sourcing_payment/constant"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type CallBack func(message []byte) error

type Handler interface {
	Handle(message []byte) error
}

type Consumer interface {
	Read(callback CallBack)
	Stop()
	SetHandler(h Handler)
	GetHandler() Handler
	GetHandlerName() string
}

type consumer struct {
	ins              *kafka.Consumer
	startOnce        sync.Once
	startClosingOnce sync.Once
	chanStop         chan bool
	producer         Producer
	dlq              bool
}

func NewConsumer(config *constant.Config) (Consumer, error) {
	return nil, nil
}
