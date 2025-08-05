package usecase

import (
	"context"
	"event_sourcing_user/application/command"
	"event_sourcing_user/constant"
	"event_sourcing_user/domain/entities"
	"event_sourcing_user/domain/value_object"
	"event_sourcing_user/infrastructure/persistent/repository"
	"event_sourcing_user/package/encrypt_password"
	"event_sourcing_user/package/ierror"
	"event_sourcing_user/package/logger"

	"go.uber.org/zap"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, command *command.CreateUser) error
}

type userUsecase struct {
	encryptPassword *encrypt_password.ArgonParam
	userRepository  repository.IUserRepository
}

func NewUserUsecase(config *constant.Config, factoryRepository repository.IRepositoryFactory) (UserUsecase, error) {
	encryptPassword, err := encrypt_password.NewArgonParam()
	if err != nil {
		return nil, ierror.Error(err)
	}
	return &userUsecase{
		encryptPassword: encryptPassword,
		userRepository:  factoryRepository.UserRepository(),
	}, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, command *command.CreateUser) error {
	log := logger.FromContext(ctx)
	email, err := value_object.NewEmail(command.Email)
	if err != nil {
		log.Error("Invalid email", zap.Error(err))
		return ierror.Error(err)
	}
	password, err := value_object.NewAndValidatePassword(command.Password)
	if err != nil {
		log.Error("Invalid password", zap.Error(err))
		return ierror.Error(err)
	}

	hashPassword, err := u.encryptPassword.HashPassword(password.Value())
	if err != nil {
		log.Error("Error hashing password", zap.Error(err))
		return ierror.Error(err)
	}

	user := entities.NewUserEntity(command.Code, email.Value(), hashPassword)
	err = u.userRepository.CreateUser(user)
	if err != nil {
		return ierror.Error(err)
	}

	return nil
}
