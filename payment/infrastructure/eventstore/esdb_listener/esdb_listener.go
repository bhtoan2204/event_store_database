package esdb_listener

import (
	"context"
	"encoding/json"
	"event_sourcing_payment/application/event/transaction"
	"event_sourcing_payment/constant"
	"event_sourcing_payment/package/eventbus"
	"event_sourcing_payment/package/logger"
	"time"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/avast/retry-go"
	"go.uber.org/zap"
)

type TransactionEventData struct {
	TransactionCode string                   `json:"transaction_code"`
	AccountID       uint                     `json:"account_id"`
	Type            constant.TransactionType `json:"type"`
	Amount          int64                    `json:"amount"`
	Reference       string                   `json:"reference"`
}

type IEventStoreListener interface {
	Start(ctx context.Context) error
}

type EsdbListener struct {
	esdbClient *esdb.Client
	eventBus   eventbus.EventBus
}

func NewEsdbListener(ctx context.Context, esdbClient *esdb.Client, eventBus eventbus.EventBus) IEventStoreListener {
	log := logger.FromContext(ctx)
	log.Info("Initializing ESDB listener")

	return &EsdbListener{
		esdbClient: esdbClient,
		eventBus:   eventBus,
	}
}

func (e *EsdbListener) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("Starting ESDB listener")

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
		retry.OnRetry(func(n uint, err error) {
			log.Warn("Retrying subscription",
				zap.Uint("attempt", n+1),
				zap.Error(err),
			)
		}),
		retry.Context(ctx),
	)

	if err != nil {
		log.Error("Subscription permanently failed after retries", zap.Error(err))
	}
}

func (e *EsdbListener) subscribe(ctx context.Context) error {
	log := logger.FromContext(ctx)

	sub, err := e.esdbClient.SubscribeToAll(ctx, esdb.SubscribeToAllOptions{
		From: esdb.Start{},
	})
	if err != nil {
		log.Error("Failed to subscribe", zap.Error(err))
		return err
	}
	defer sub.Close()

	log.Info("Successfully subscribed to EventStore")

	// only allow 100 concurrent events to be processed at a time
	concurrencyLimit := 100
	semaphore := make(chan struct{}, concurrencyLimit)

	for {
		select {
		case <-ctx.Done():
			log.Info("Listener stopped via context cancel")
			return nil
		default:
			subEvent := sub.Recv()
			if subEvent == nil {
				continue
			}

			if subEvent.SubscriptionDropped != nil {
				log.Error("Subscription dropped", zap.Error(subEvent.SubscriptionDropped.Error))
				return subEvent.SubscriptionDropped.Error
			}

			if subEvent.CheckPointReached != nil || subEvent.CaughtUp != nil || subEvent.FellBehind != nil {
				continue
			}

			if subEvent.EventAppeared != nil {
				original := subEvent.EventAppeared.OriginalEvent()
				if original == nil || original.EventType == "" {
					continue
				}

				log.Debug("Received event",
					zap.String("stream", original.StreamID),
					zap.String("eventType", original.EventType),
				)

				semaphore <- struct{}{}

				go func(ev *esdb.RecordedEvent) {
					defer func() {
						<-semaphore
					}()

					switch original.EventType {
					case "TransactionDeposit", "TransactionWithdraw", "TransactionTransfer":
						var txEvent TransactionEventData
						if err := json.Unmarshal(original.Data, &txEvent); err != nil {
							log.Error("Failed to unmarshal transaction event",
								zap.Error(err),
								zap.String("eventType", original.EventType))
							return
						}

						createdEvent := transaction.NewTransactionCreatedEvent(
							txEvent.AccountID,
							txEvent.Amount,
							txEvent.Type,
							txEvent.TransactionCode,
						)

						if err := e.eventBus.Dispatch(ctx, createdEvent); err != nil {
							log.Error("Failed to dispatch transaction event", zap.Error(err))
						}

					default:
						log.Debug("Unhandled event type", zap.String("eventType", original.EventType))
					}
				}(original)
			}
		}
	}
}
