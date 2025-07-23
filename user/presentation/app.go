package presentation

import (
	"context"
	"event_sourcing_user/constant"
	"event_sourcing_user/domain/usecase"
	"event_sourcing_user/infrastructure/persistent"
	"event_sourcing_user/infrastructure/persistent/persistent_object"
	"event_sourcing_user/infrastructure/persistent/repository"
	"event_sourcing_user/infrastructure/srvdisc"
	"event_sourcing_user/package/grpc_infra"
	"event_sourcing_user/package/logger"
	"event_sourcing_user/package/server"
	"event_sourcing_user/presentation/grpc_layer"
	"event_sourcing_user/proto/user"

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
	cfg *constant.Config

	repositoryFactory repository.IRepositoryFactory
	usecase           usecase.Usecase
}

func NewApp(ctx context.Context, cfg *constant.Config) (App, error) {
	log := logger.FromContext(ctx)
	persistentConnection, err := persistent.NewPersistentConnection(&cfg.Postgres)
	if err != nil {
		log.Error("Error initializing persistent connection", zap.Error(err))
		return nil, err
	}
	err = persistentConnection.SyncTable(&persistent_object.User{})
	if err != nil {
		log.Error("Error syncing table", zap.Error(err))
		return nil, err
	}

	repositoryFactory := repository.NewRepositoryFactory(persistentConnection)
	usecase, err := usecase.NewUsecase(cfg, repositoryFactory)
	if err != nil {
		log.Error("Error creating usecase", zap.Error(err))
		return nil, err
	}
	return &app{cfg: cfg, repositoryFactory: repositoryFactory, usecase: usecase}, nil
}

func (a *app) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Info("Starting server")
	panicHandler := func(p any) (err error) {
		return status.Errorf(codes.Internal, "%s", p)
	}

	var sopts []grpc.ServerOption
	sopts = append(sopts,
		grpc.ChainUnaryInterceptor(
			grpc_infra.MonitorRequestDuration(nil),
			grpc_infra.Recovery(panicHandler),
			grpc_infra.Timeout(),
			grpc_infra.HandleError(),
		),
	)
	rpcServer := grpc.NewServer(sopts...)

	healthCheck := grpc_infra.NewHealthService()
	grpc_health_v1.RegisterHealthServer(rpcServer, healthCheck)

	userService, err := grpc_layer.NewGrpcPresentation(rpcServer, a.repositoryFactory, a.cfg, a.usecase)
	if err != nil {
		log.Error("Error creating gRPC presentation", zap.Error(err))
		return err
	}
	user.RegisterUserServiceServer(rpcServer, userService)

	grpcServer, err := server.New()
	if err != nil {
		log.Error("Error creating gRPC server", zap.Error(err))
		return err
	}

	consulConnection, err := srvdisc.NewConsulServiceDiscovery(ctx, a.cfg)
	if err != nil {
		log.Error("Error creating Consul service discovery", zap.Error(err))
		return err
	}

	consulConnection.Register(ctx, a.cfg.Server.ServiceName, grpcServer.PortInt())
	defer consulConnection.DeRegister(ctx, a.cfg.Server.ServiceName)

	return grpcServer.ServeGRPC(ctx, rpcServer)
}
