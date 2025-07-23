package usecase

import (
	"event_sourcing_user/constant"
	"event_sourcing_user/infrastructure/persistent/repository"
	"event_sourcing_user/package/ierror"
)

type Usecase interface {
	UserUsecase() UserUsecase
	AuthUsecase() AuthUsecase
}

type usecase struct {
	config      *constant.Config
	userUsecase UserUsecase
	authUsecase AuthUsecase
}

func NewUsecase(config *constant.Config, factoryRepository repository.IRepositoryFactory) (Usecase, error) {
	userUsecase, err := NewUserUsecase(config, factoryRepository)
	if err != nil {
		return nil, ierror.Error(err)
	}
	authUsecase, err := NewAuthUsecase(config, factoryRepository)
	if err != nil {
		return nil, ierror.Error(err)
	}
	return &usecase{
		config:      config,
		userUsecase: userUsecase,
		authUsecase: authUsecase,
	}, nil
}

func (u *usecase) UserUsecase() UserUsecase {
	return u.userUsecase
}

func (u *usecase) AuthUsecase() AuthUsecase {
	return u.authUsecase
}
