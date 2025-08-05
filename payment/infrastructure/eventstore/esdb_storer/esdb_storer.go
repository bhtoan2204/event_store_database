package esdb_storer

import (
	"context"
	"encoding/json"
	"event_sourcing_payment/application/event"
	"event_sourcing_payment/package/logger"
	"fmt"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type IEventStorer interface {
	Append(ctx context.Context, aggregateID string, events []event.DomainEvent) error
}

type EsdbStorer struct {
	esdbClient *esdb.Client
}

func NewEsdbStorer(ctx context.Context, esdbClient *esdb.Client) IEventStorer {
	return &EsdbStorer{esdbClient: esdbClient}
}

func (s *EsdbStorer) Append(ctx context.Context, aggregateID string, events []event.DomainEvent) error {
	log := logger.FromContext(ctx)

	eventDataList := make([]esdb.EventData, 0, len(events))
	for _, evt := range events {
		data, err := json.Marshal(evt)
		if err != nil {
			log.Error("Failed to marshal event", zap.Error(err))
			return err
		}
		eventDataList = append(eventDataList, esdb.EventData{
			ContentType: esdb.ContentTypeJson,
			EventType:   evt.EventName(),
			Data:        data,
			EventID:     uuid.Must(uuid.NewV4()),
		})
	}

	streamID := fmt.Sprintf("account-%s", aggregateID)
	_, err := s.esdbClient.AppendToStream(ctx, streamID, esdb.AppendToStreamOptions{}, eventDataList...)
	return err
}
