package esdb_listener

import (
	"context"
	"encoding/json"
	"event_sourcing_payment/application/event"
	"event_sourcing_payment/package/eventbus"
	"event_sourcing_payment/package/logger"
	"time"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/avast/retry-go"
	"go.uber.org/zap"
)

type EsdbListener struct {
	esdbClient *esdb.Client
	eventBus   eventbus.EventBus
}

func NewEsdbListener(esdbClient *esdb.Client, eventBus eventbus.EventBus) *EsdbListener {
	return &EsdbListener{esdbClient: esdbClient, eventBus: eventBus}
}

func (e *EsdbListener) Start(ctx context.Context) error {
	go e.runSubscription(ctx)
	return nil
}

func (e *EsdbListener) runSubscription(ctx context.Context) {
	log := logger.FromContext(ctx)

	err := retry.Do(
		func() error {
			log.Info("Attempting to subscribe to ESDB")
			return e.subscribe(ctx)
		},
		retry.Attempts(5),
		retry.Delay(5*time.Second),
		retry.DelayType(retry.FixedDelay),
		retry.LastErrorOnly(true),
		retry.Context(ctx),
	)

	if err != nil {
		log.Error("Subscription permanently failed after retries", zap.Error(err))
	}
}

func (e *EsdbListener) subscribe(ctx context.Context) error {
	log := logger.FromContext(ctx)

	sub, err := e.esdbClient.SubscribeToAll(ctx, esdb.SubscribeToAllOptions{From: esdb.Start{}})
	if err != nil {
		log.Error("Failed to subscribe", zap.Error(err))
		return err
	}
	defer sub.Close()

	semaphore := make(chan struct{}, 100)

	for {
		select {
		case <-ctx.Done():
			log.Info("Listener stopped via context cancel")
			return nil
		default:
			subEvent := sub.Recv()
			if subEvent == nil || subEvent.EventAppeared == nil {
				continue
			}

			original := subEvent.EventAppeared.OriginalEvent()
			if original == nil || original.EventType == "" {
				continue
			}

			log.Debug("Received event", zap.String("stream", original.StreamID), zap.String("eventType", original.EventType))

			semaphore <- struct{}{}

			go func(ev *esdb.RecordedEvent) {
				defer func() { <-semaphore }()

				var evt event.TransactionCreatedEvent
				if err := json.Unmarshal(ev.Data, &evt); err != nil {
					log.Error("Failed to unmarshal event", zap.Error(err))
					return
				}

				if err := e.eventBus.Dispatch(ctx, evt); err != nil {
					log.Error("Failed to dispatch event", zap.Error(err))
				}
			}(original)
		}
	}
}
