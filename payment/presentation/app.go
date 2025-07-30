package presentation

import (
	"context"
	"event_sourcing_payment/constant"
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
	cfg *constant.Config
}

func NewApp(ctx context.Context, cfg *constant.Config) (App, error) {
	// log := logger.FromContext(ctx)

	return &app{cfg: cfg}, nil
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

	paymentService, err := grpc_layer.NewGrpcPresentation(rpcServer, a.cfg)
	if err != nil {
		log.Error("Error creating gRPC presentation", zap.Error(err))
		return err
	}
	payment.RegisterPaymentServiceServer(rpcServer, paymentService)

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
