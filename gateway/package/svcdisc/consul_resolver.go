package svcdisc

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
)

type consulResolver struct {
	consulClient *api.Client
	serviceName  string

	cc     resolver.ClientConn
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func (c *Consul) NewResolverBuilder(serviceName string) resolver.Builder {
	return &consulResolver{
		consulClient: c.consulClient,
		serviceName:  serviceName,
	}
}

func (r *consulResolver) Scheme() string { return "consul" }

func (r *consulResolver) Build(target resolver.Target, cc resolver.ClientConn, _ resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc
	r.ctx, r.cancel = context.WithCancel(context.Background())

	r.wg.Add(1)
	go r.watch()

	// trigger first resolve immediately
	addrs, err := r.getAddresses()
	if err == nil {
		r.updateState(addrs)
	}

	return r, nil
}

func (r *consulResolver) ResolveNow(resolver.ResolveNowOptions) {}

func (r *consulResolver) Close() {
	r.cancel()
	r.wg.Wait()
}

func (r *consulResolver) watch() {
	defer r.wg.Done()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-r.ctx.Done():
			return
		case <-ticker.C:
			addrs, err := r.getAddresses()
			if err != nil {
				// log your error here
				continue
			}
			r.updateState(addrs)
		}
	}
}

func (r *consulResolver) getAddresses() ([]resolver.Address, error) {
	entries, _, err := r.consulClient.Health().Service(r.serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	addresses := make([]resolver.Address, 0, len(entries))
	for _, entry := range entries {
		svc := entry.Service
		host := svc.Address
		if host == "" {
			host = entry.Node.Address
		}
		// strip port if present in Address
		if strings.Contains(host, ":") {
			h, _, e := net.SplitHostPort(host)
			if e == nil {
				host = h
			}
		}
		addresses = append(addresses, resolver.Address{
			Addr: fmt.Sprintf("%s:%d", host, svc.Port),
		})
	}
	return addresses, nil
}

func (r *consulResolver) updateState(addrs []resolver.Address) {
	_ = r.cc.UpdateState(resolver.State{Addresses: addrs})
}
