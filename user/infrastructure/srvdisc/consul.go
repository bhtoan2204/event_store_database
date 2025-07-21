package srvdisc

import (
	"event_sourcing_user/constant"
	"fmt"

	"github.com/hashicorp/consul/api"
)

type ServiceDiscovery interface {
	Register(serviceName string, servicePort int) error
	DeRegister(serviceName string) error
}

type ConsulServiceDiscovery struct {
	consulClient *api.Client
	rrCounter    uint64
}

func NewConsulServiceDiscovery(cfg *constant.Config) ServiceDiscovery {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = cfg.ConsulConfig.Address
	consulConfig.Scheme = cfg.ConsulConfig.Scheme
	consulConfig.Datacenter = cfg.ConsulConfig.DataCenter
	consulConfig.Token = cfg.ConsulConfig.Token

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}

	return &ConsulServiceDiscovery{
		consulClient: consulClient,
		rrCounter:    0,
	}
}

func (s *ConsulServiceDiscovery) Register(serviceName string, servicePort int) error {
	serviceID := fmt.Sprintf("%s-%d", serviceName, s.rrCounter)
	s.rrCounter++

	registration := &api.AgentServiceRegistration{
		ID:    serviceID,
		Name:  serviceName,
		Port:  servicePort,
		Check: &api.AgentServiceCheck{},
	}

	if err := s.consulClient.Agent().ServiceRegister(registration); err != nil {
		return err
	}

	return nil
}

func (s *ConsulServiceDiscovery) DeRegister(serviceName string) error {
	serviceID := fmt.Sprintf("%s-%d", serviceName, s.rrCounter)
	s.rrCounter++

	if err := s.consulClient.Agent().ServiceDeregister(serviceID); err != nil {
		return err
	}

	return nil
}
