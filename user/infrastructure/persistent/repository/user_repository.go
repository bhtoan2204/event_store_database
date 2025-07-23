package repository

import (
	"event_sourcing_user/domain/entities"
	"event_sourcing_user/infrastructure/persistent/mapper"
	"event_sourcing_user/infrastructure/persistent/persistent_object"

	"gorm.io/gorm"
)

var _ IUserRepository = &userRepository{}

type IUserRepository interface {
	CreateUser(user *entities.UserEntity) error
	GetUserByID(id int64) (*entities.UserEntity, error)
	GetUserByEmail(email string) (*entities.UserEntity, error)
	UpdateUser(user *entities.UserEntity) error
	DeleteUser(id int64) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *entities.UserEntity) error {
	persistentUser := mapper.UserEntityToUser(user)
	if err := r.db.Create(persistentUser).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUserByID(id int64) (*entities.UserEntity, error) {
	var user persistent_object.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	entityUser := mapper.UserToUserEntity(&user)
	return entityUser, nil
}

func (r *userRepository) GetUserByEmail(email string) (*entities.UserEntity, error) {
	var user persistent_object.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	entityUser := mapper.UserToUserEntity(&user)
	return entityUser, nil
}

func (r *userRepository) UpdateUser(user *entities.UserEntity) error {
	persistentUser := mapper.UserEntityToUser(user)
	return r.db.Save(persistentUser).Error
}

func (r *userRepository) DeleteUser(id int64) error {
	return r.db.Delete(&persistent_object.User{}, id).Error
}
