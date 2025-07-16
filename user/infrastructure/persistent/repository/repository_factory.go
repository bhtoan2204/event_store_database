package repository

import (
	"context"

	"gorm.io/gorm"
)

type IRepositoryFactory interface {
	WithTransaction(ctx context.Context, fn func(IRepositoryFactory) error) error
	UserRepository() IUserRepository
}

type repositoryFactory struct {
	db             *gorm.DB
	userRepository IUserRepository
}

func NewRepositoryFactory(db *gorm.DB) IRepositoryFactory {
	userRepository := NewUserRepository(db)
	return &repositoryFactory{db: db, userRepository: userRepository}
}

func (r *repositoryFactory) WithTransaction(ctx context.Context, fn func(IRepositoryFactory) error) (err error) {
	tx := r.db.Begin()
	tr := NewRepositoryFactory(tx)

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
