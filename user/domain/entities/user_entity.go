package entities

import "time"

type UserEntity struct {
	id        int64
	code      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
}

func NewUserEntity(code string, email string, password string) *UserEntity {
	return &UserEntity{
		code:     code,
		email:    email,
		password: password,
	}
}

func (u *UserEntity) ID() int64 {
	return u.id
}

func (u *UserEntity) Code() string {
	return u.code
}

func (u *UserEntity) Email() string {
	return u.email
}

func (u *UserEntity) Password() string {
	return u.password
}

func (u *UserEntity) CreatedAt() time.Time {
	return u.createdAt
}

func (u *UserEntity) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u *UserEntity) SetID(id int64) {
	u.id = id
}

func (u *UserEntity) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *UserEntity) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}
