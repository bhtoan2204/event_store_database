package srvdisc

import (
	"context"
	"event_sourcing_user/constant"
	"event_sourcing_user/package/logger"
	"event_sourcing_user/utils"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type ServiceDiscovery interface {
	Register(ctx context.Context, serviceName string, servicePort int) error
	DeRegister(ctx context.Context, serviceName string) error
}

type ConsulServiceDiscovery struct {
	consulClient *api.Client
	id           string
	rrCounter    uint64
}

func NewConsulServiceDiscovery(ctx context.Context, cfg *constant.Config) (ServiceDiscovery, error) {
	log := logger.FromContext(ctx)
	consulConfig := api.DefaultConfig()
	consulConfig.Address = cfg.ConsulConfig.Address
	consulConfig.Scheme = cfg.ConsulConfig.Scheme
	consulConfig.Datacenter = cfg.ConsulConfig.DataCenter
	consulConfig.Token = cfg.ConsulConfig.Token

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		log.Error("Error creating Consul client", zap.Error(err))
		return nil, err
	}

	return &ConsulServiceDiscovery{
		consulClient: consulClient,
		rrCounter:    0,
	}, nil
}

func (s *ConsulServiceDiscovery) Register(ctx context.Context, serviceName string, servicePort int) error {
	log := logger.FromContext(ctx)
	serviceID := uuid.New().String()
	serviceAddress, err := utils.GetInternalIP()
	if err != nil {
		return err
	}
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    serviceName,
		Port:    servicePort,
		Address: serviceAddress,
		Tags:    []string{"api", "user"},
		Check: &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", serviceAddress, servicePort),
			Interval:                       "15s",
			Timeout:                        "5s",
			DeregisterCriticalServiceAfter: "30s",
		},
	}

	if err := s.consulClient.Agent().ServiceRegister(registration); err != nil {
		log.Error("Error registering service", zap.Error(err))
		return err
	}
	s.id = serviceID

	return nil
}

func (s *ConsulServiceDiscovery) DeRegister(ctx context.Context, serviceName string) error {
	log := logger.FromContext(ctx)
	if err := s.consulClient.Agent().ServiceDeregister(s.id); err != nil {
		return err
	}
	log.Info("DeRegister service successfully", zap.String("serviceID", s.id))
	return nil
}
