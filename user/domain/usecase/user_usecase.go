package usecase

import (
	"event_sourcing_user/application/command"
	"event_sourcing_user/domain/aggregation"
	"event_sourcing_user/domain/entities"
	"event_sourcing_user/domain/value_object"
	"event_sourcing_user/infrastructure/persistent/repository"
	"log"
)

var _ UserUC = (*userUC)(nil)

type UserUC interface {
	CreateUser(command *command.CreateUserCommand) error
	GetUserByID(id int64) (*aggregation.UserAggregation, error)
	GetUserByEmail(email string) (*aggregation.UserAggregation, error)
	UpdateUser(user *entities.UserEntity) error
	DeleteUser(id int64) error
}

type userUC struct {
	userRepository repository.IUserRepository
}

func NewUserUC(userRepository repository.IUserRepository) UserUC {
	return &userUC{userRepository: userRepository}
}

func (uc *userUC) CreateUser(command *command.CreateUserCommand) error {
	passwordObject, err := value_object.NewAndValidatePassword(command.Password)
	if err != nil {
		log.Fatal(err)
		return err
	}
	user := entities.NewUserEntity(command.Email, passwordObject.Value())
	if err := uc.userRepository.CreateUser(user); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (uc *userUC) GetUserByID(id int64) (*aggregation.UserAggregation, error) {
	user, err := uc.userRepository.GetUserByID(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return aggregation.NewUserAggregation(user), nil
}

func (uc *userUC) GetUserByEmail(email string) (*aggregation.UserAggregation, error) {
	user, err := uc.userRepository.GetUserByEmail(email)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return aggregation.NewUserAggregation(user), nil
}

func (uc *userUC) UpdateUser(user *entities.UserEntity) error {
	if err := uc.userRepository.UpdateUser(user); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (uc *userUC) DeleteUser(id int64) error {
	if err := uc.userRepository.DeleteUser(id); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
