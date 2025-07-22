package svcdisc

import (
	"context"
	"event_sourcing_gateway/package/logger"
	"event_sourcing_gateway/package/settings"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type Consul struct {
	consulClient *api.Client
}

func NewConsul(config *settings.Config) *Consul {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = config.ConsulConfig.Address
	consulConfig.Scheme = config.ConsulConfig.Scheme
	consulConfig.Datacenter = config.ConsulConfig.DataCenter
	consulConfig.Token = config.ConsulConfig.Token

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}

	return &Consul{consulClient: consulClient}
}

func (c *Consul) GetService(ctx context.Context, serviceName string) (*api.AgentService, error) {
	log := logger.FromContext(ctx)
	services, err := c.consulClient.Agent().Services()
	if err != nil {
		log.Error("failed to get services", zap.Error(err))
		return nil, err
	}

	return services[serviceName], nil
}
