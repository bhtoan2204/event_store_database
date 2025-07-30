package messaging_layer

import (
	"event_sourcing_payment/constant"
	"event_sourcing_payment/infrastructure/broker"
)

type Messaging interface {
	Consumer() broker.Consumer
}

type messaging struct {
	consumer broker.Consumer
}

func NewMessaging(config *constant.Config) (Messaging, error) {
	consumer, err := broker.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return &messaging{consumer: consumer}, nil
}

func (m *messaging) Consumer() broker.Consumer {
	return m.consumer
}
