package eventstore

import (
	"context"
	"event_sourcing_payment/constant"
	"event_sourcing_payment/package/logger"
	"strconv"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"go.uber.org/zap"
)

var _ IEventStoreConnection = &EventStoreConnection{}

type IEventStoreConnection interface {
	GetClient() *esdb.Client
}

type EventStoreConnection struct {
	esdbClient *esdb.Client
}

func NewEventStoreConnection(ctx context.Context, config *constant.Config) (IEventStoreConnection, error) {
	log := logger.FromContext(ctx)
	connectionString := "esdb://" + config.EventStore.Host + ":" + strconv.Itoa(config.EventStore.Port) + "?tls=false"
	settings, err := esdb.ParseConnectionString(connectionString)
	if err != nil {
		log.Error("Error parsing connection string", zap.Error(err))
		return nil, err
	}
	db, err := esdb.NewClient(settings)
	if err != nil {
		log.Error("Error creating ESDB client", zap.Error(err))
		return nil, err
	}

	return &EventStoreConnection{
		esdbClient: db,
	}, nil
}

func (e *EventStoreConnection) GetClient() *esdb.Client {
	return e.esdbClient
}
