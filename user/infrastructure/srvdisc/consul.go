package srvdisc

import (
	"event_sourcing_user/constant"
	"event_sourcing_user/utils"
	"fmt"
	"log"

	"github.com/google/uuid"
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

func NewConsulServiceDiscovery(cfg *constant.Config) (ServiceDiscovery, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = cfg.ConsulConfig.Address
	consulConfig.Scheme = cfg.ConsulConfig.Scheme
	consulConfig.Datacenter = cfg.ConsulConfig.DataCenter
	consulConfig.Token = cfg.ConsulConfig.Token

	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return nil, err
	}

	return &ConsulServiceDiscovery{
		consulClient: consulClient,
		rrCounter:    0,
	}, nil
}

func (s *ConsulServiceDiscovery) Register(serviceName string, servicePort int) error {
	serviceID := uuid.New().String()
	serviceAddress, err := utils.GetInternalIP()
	if err != nil {
		return err
	}
	s.rrCounter++
	fmt.Println("serviceID", serviceID)
	fmt.Println("serviceName", serviceName)
	fmt.Println("serviceAddress", serviceAddress)
	fmt.Println("servicePort", servicePort)
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
		log.Println("Error registering service", err)
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
