package aggregation

import "event_sourcing_user/domain/entities"

type UserAggregation struct {
	user *entities.UserEntity
}

func NewUserAggregation(user *entities.UserEntity) *UserAggregation {
	return &UserAggregation{
		user: user,
	}
}
