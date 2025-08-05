package command

import (
	"errors"
	"event_sourcing_user/proto/user"

	"github.com/google/uuid"
)

type CreateUser struct {
	Code     string
	Email    string
	Password string
}

func NewCreateUser(req *user.CreateUserRequest) *CreateUser {
	code := uuid.New().String()
	return &CreateUser{
		Code:     code,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (c *CreateUser) Validate() error {
	if c.Code == "" {
		return errors.New("code is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	if c.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
