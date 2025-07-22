package presentation

import (
	"context"
	"event_sourcing_user/constant"
	"event_sourcing_user/infrastructure/persistent"
	"event_sourcing_user/infrastructure/persistent/repository"
	"event_sourcing_user/infrastructure/srvdisc"
	"event_sourcing_user/package/grpc_infra"
	"event_sourcing_user/package/server"
	"event_sourcing_user/presentation/grpc_layer"
	"event_sourcing_user/proto/user"
	"log"

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
}

func NewApp(cfg *constant.Config) (App, error) {
	persistentConnection, err := persistent.NewPersistentConnection(&cfg.Postgres)
	if err != nil {
		return nil, err
	}

	repositoryFactory := repository.NewRepositoryFactory(persistentConnection)
	return &app{cfg: cfg, repositoryFactory: repositoryFactory}, nil
}

func (a *app) Start(ctx context.Context) error {
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

	userService := grpc_layer.NewGrpcPresentation(rpcServer, a.repositoryFactory)
	user.RegisterUserServiceServer(rpcServer, userService)

	grpcServer, err := server.New()
	if err != nil {
		log.Println("Error creating gRPC server", err)
		return err
	}

	consulConnection, err := srvdisc.NewConsulServiceDiscovery(a.cfg)
	consulConnection.Register(a.cfg.Server.ServiceName, grpcServer.PortInt())
	if err != nil {
		log.Println("Error creating Consul service discovery", err)
		return err
	}

	log.Println("Starting server")
	return grpcServer.ServeGRPC(ctx, rpcServer)
}
