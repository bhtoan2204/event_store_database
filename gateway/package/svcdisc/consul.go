package svcdisc

import (
	"context"
	"event_sourcing_gateway/package/logger"
	"event_sourcing_gateway/package/settings"
	"fmt"
	"sync/atomic"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type Consul struct {
	consulClient *api.Client
	rr           uint64
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

	return &Consul{consulClient: consulClient, rr: 0}
}

func (c *Consul) GetService(ctx context.Context, serviceName string) (*api.AgentService, error) {
	if serviceName == "" {
		return nil, fmt.Errorf("service name is required")
	}

	log := logger.FromContext(ctx)

	entries, _, err := c.consulClient.Health().Service(serviceName, "", true, nil)
	if err != nil {
		log.Error("failed to query health service", zap.Error(err))
		return nil, err
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("service %s not found or not healthy", serviceName)
	}

	i := atomic.AddUint64(&c.rr, 1) - 1
	return entries[i%uint64(len(entries))].Service, nil
}
