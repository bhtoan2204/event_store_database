package presentation

import (
	"context"

	"event_sourcing_payment/application/command"
	"event_sourcing_payment/application/event"
	"event_sourcing_payment/application/query"
	"event_sourcing_payment/constant"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/infrastructure/adapter"
	"event_sourcing_payment/infrastructure/eventstore"
	"event_sourcing_payment/infrastructure/eventstore/esdb_listener"
	"event_sourcing_payment/infrastructure/eventstore/esdb_storer"
	"event_sourcing_payment/infrastructure/projection"
	"event_sourcing_payment/infrastructure/projection/persistent_object"
	"event_sourcing_payment/infrastructure/projection/repository"
	"event_sourcing_payment/infrastructure/srvdisc"
	"event_sourcing_payment/package/grpc_infra"
	"event_sourcing_payment/package/logger"
	"event_sourcing_payment/package/server"
	"event_sourcing_payment/presentation/grpc_layer"
	"event_sourcing_payment/proto/payment"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type App interface {
	Start(ctx context.Context) error
}

type app struct {
	cfg         *constant.Config
	adapter     adapter.IAdapter
	commandBus  *command.CommandBus
	eventBus    *event.EventBus
	queryBus    *query.QueryBus
	eventStorer esdb_storer.IEventStorer
	useCase     usecase.IUseCase
}

func NewApp(ctx context.Context, cfg *constant.Config) (App, error) {
	log := logger.FromContext(ctx)
	log.Info("Initializing app")

	// Connect to Projection DB
	projectionConn, err := projection.NewProjectionConnection(ctx, &cfg.Postgres)
	if err != nil {
		log.Error("Failed to connect to projection DB", zap.Error(err))
		return nil, err
	}
	err = projectionConn.SyncTable(
		&persistent_object.Account{},
		&persistent_object.Transaction{},
		&persistent_object.Outbox{},
	)
	if err != nil {
		log.Error("Failed to sync projection tables", zap.Error(err))
		return nil, err
	}

	// Connect to Event Store
	eventStoreConn, err := eventstore.NewEventStoreConnection(ctx, cfg)
	if err != nil {
		log.Error("Failed to connect to event store", zap.Error(err))
		return nil, err
	}

	// Redis and Locker
	adapter, err := adapter.NewAdapter(ctx, cfg)
	if err != nil {
		log.Error("Failed to initialize adapter", zap.Error(err))
		return nil, err
	}

	repo := repository.NewFactoryRepository(ctx, projectionConn)
	esdbStorer := esdb_storer.NewEsdbStorer(ctx, eventStoreConn.GetClient())
	useCase := usecase.NewUseCase(esdbStorer, repo)
	eventBus := event.NewEventBus()
	commandBus := command.NewCommandBus(esdbStorer, useCase)
	queryBus := query.NewQueryBus(repo, useCase)

	// Start event listener
	esdbListener := esdb_listener.NewEsdbListener(eventStoreConn.GetClient(), eventBus)
	esdbListener.Start(ctx)

	log.Info("App initialized successfully")

	return &app{
		cfg:         cfg,
		adapter:     adapter,
		commandBus:  commandBus,
		eventBus:    eventBus,
		queryBus:    queryBus,
		eventStorer: esdbStorer,
		useCase:     useCase,
	}, nil
}

func (a *app) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("Starting gRPC server")

	panicHandler := func(p any) (err error) {
		return status.Errorf(codes.Internal, "%s", p)
	}

	// gRPC interceptors
	var serverOpts []grpc.ServerOption
	serverOpts = append(serverOpts,
		grpc.ChainUnaryInterceptor(
			grpc_infra.MonitorRequestDuration(nil),
			grpc_infra.Recovery(panicHandler),
			grpc_infra.Timeout(),
			grpc_infra.HandleError(),
		),
	)

	rpcServer := grpc.NewServer(serverOpts...)

	// Health check service
	grpc_health_v1.RegisterHealthServer(rpcServer, grpc_infra.NewHealthService())

	// Register payment service
	paymentService, err := grpc_layer.NewGrpcPresentation(rpcServer, a.commandBus)
	if err != nil {
		log.Error("Failed to initialize payment service", zap.Error(err))
		return err
	}
	payment.RegisterPaymentServiceServer(rpcServer, paymentService)

	// Create GRPC Server
	grpcSrv, err := server.New()
	if err != nil {
		log.Error("Failed to create GRPC server", zap.Error(err))
		return err
	}

	// Consul service registration
	consulConn, err := srvdisc.NewConsulServiceDiscovery(ctx, a.cfg)
	if err != nil {
		log.Error("Failed to initialize Consul", zap.Error(err))
		return err
	}
	consulConn.Register(ctx, a.cfg.Server.ServiceName, grpcSrv.PortInt())
	defer consulConn.DeRegister(ctx, a.cfg.Server.ServiceName)

	return grpcSrv.ServeGRPC(ctx, rpcServer)
}
