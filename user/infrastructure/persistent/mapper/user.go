package mapper

import (
	"event_sourcing_user/domain/entities"
	"event_sourcing_user/infrastructure/persistent/persistent_object"
)

func UserEntityToUser(user *entities.UserEntity) *persistent_object.User {
	persistentUser := persistent_object.NewUser(user.Code(), user.Email(), user.Password())
	return persistentUser
}

func UserToUserEntity(user *persistent_object.User) *entities.UserEntity {
	entityUser := entities.NewUserEntity(user.Code, user.Email, user.Password)
	entityUser.SetID(user.Base.ID)
	entityUser.SetCreatedAt(user.CreatedAt)
	entityUser.SetUpdatedAt(user.UpdatedAt)
	return entityUser
}
