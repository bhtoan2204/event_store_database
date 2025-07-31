package esdb_storer

import (
	"context"
	"encoding/json"
	event "event_sourcing_payment/application/event/transaction"
	"event_sourcing_payment/package/logger"
	"fmt"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type IEventStorer interface {
	SaveTransactionEvent(ctx context.Context, txEvent *event.TransactionCreatedEvent) error
}

type EsdbStorer struct {
	esdbClient *esdb.Client
}

func NewEsdbStorer(ctx context.Context, esdbClient *esdb.Client) IEventStorer {
	log := logger.FromContext(ctx)
	log.Info("Initializing ESDB storer")

	return &EsdbStorer{esdbClient: esdbClient}
}

func (e *EsdbStorer) SaveTransactionEvent(ctx context.Context, txEvent *event.TransactionCreatedEvent) error {
	log := logger.FromContext(ctx)

	data, err := json.Marshal(txEvent)
	if err != nil {
		log.Error("Failed to marshal transaction event", zap.Error(err))
		return err
	}

	metadata, _ := json.Marshal(txEvent)

	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   txEvent.Type.String(),
		Data:        data,
		Metadata:    metadata,
		EventID:     uuid.Must(uuid.NewV4()),
	}

	// Best practice: use account ID as stream ID
	streamID := fmt.Sprintf("account-%d", txEvent.AccountID)

	_, err = e.esdbClient.AppendToStream(ctx, streamID, esdb.AppendToStreamOptions{}, eventData)
	if err != nil {
		log.Error("Failed to append transaction event to EventStore", zap.Error(err))
		return err
	}

	return nil
}
