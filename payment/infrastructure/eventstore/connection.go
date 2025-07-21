package eventstore

import (
	"log"

	esdb "github.com/EventStore/EventStore-Client-Go/v3/esdb"
)

var _ IEventStoreConnection = &EventStoreConnection{}

type IEventStoreConnection interface {
	GetClient() *esdb.Client
}

type EventStoreConnection struct {
	esdbClient *esdb.Client
}

func NewEventStoreConnection() IEventStoreConnection {
	settings, err := esdb.ParseConnectionString("esdb://localhost:2113?tls=false")
	if err != nil {
		log.Fatalf("Lỗi cấu hình ESDB: %v", err)
	}

	db, err := esdb.NewClient(settings)
	if err != nil {
		log.Fatalf("Lỗi khởi tạo ESDB client: %v", err)
	}

	return &EventStoreConnection{
		esdbClient: db,
	}
}

func (e *EventStoreConnection) GetClient() *esdb.Client {
	return e.esdbClient
}
