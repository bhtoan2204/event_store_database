package command

import (
	"errors"
	"event_sourcing_user/proto/user"
)

type LoginCommand struct {
	Email    string
	Password string
}

func NewLoginCommand(req *user.LoginRequest) *LoginCommand {
	return &LoginCommand{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (c *LoginCommand) Validate() error {
	if c.Email == "" {
		return errors.New("email is required")
	}
	if c.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
