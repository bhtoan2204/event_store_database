package routing

import "event_sourcing_gateway/application/model"

type RoutingUseCase interface {
	Forward(routingData *model.RoutingData) (interface{}, error)
}
