package routing

import "event_sourcing_gateway/application/model"

type ServiceClient interface {
	Invoke(routingData *model.RoutingData) (interface{}, error)
}
