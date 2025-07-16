package repository

import (
	"event_sourcing_user/infrastructure/persistent/persistent_object"

	"gorm.io/gorm"
)

var _ IUserRepository = &userRepository{}

type IUserRepository interface {
	CreateUser(user *persistent_object.User) error
	GetUserByID(id string) (*persistent_object.User, error)
	GetUserByEmail(email string) (*persistent_object.User, error)
	UpdateUser(user *persistent_object.User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *persistent_object.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id string) (*persistent_object.User, error) {
	return nil, nil
}

func (r *userRepository) GetUserByEmail(email string) (*persistent_object.User, error) {
	var user persistent_object.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *persistent_object.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&persistent_object.User{}, id).Error
}
