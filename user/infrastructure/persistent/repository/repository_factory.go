package repository

import (
	"context"
	"event_sourcing_user/infrastructure/persistent"
)

type IRepositoryFactory interface {
	WithTransaction(ctx context.Context, fn func(IRepositoryFactory) error) error
	UserRepository() IUserRepository
}

type repositoryFactory struct {
	persistentConnection *persistent.PersistentConnection
	userRepository       IUserRepository
}

func NewRepositoryFactory(persistentConnection *persistent.PersistentConnection) IRepositoryFactory {
	userRepository := NewUserRepository(persistentConnection.GetDB())
	return &repositoryFactory{persistentConnection: persistentConnection, userRepository: userRepository}
}

func (r *repositoryFactory) WithTransaction(ctx context.Context, fn func(IRepositoryFactory) error) (err error) {
	tx := r.persistentConnection.GetDB().Begin()
	tr := NewRepositoryFactory(r.persistentConnection)

	err = tx.Error
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	err = fn(tr)

	return err
}

func (f *repositoryFactory) UserRepository() IUserRepository {
	return f.userRepository
}
