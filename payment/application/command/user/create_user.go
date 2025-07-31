package command

import (
	"context"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/package/commandbus"
	"fmt"
)

type CreateUserCommand struct {
	Username string
	Email    string
}

func (c CreateUserCommand) CommandName() string {
	return "CreateUserCommand"
}

type CreateUserHandler struct {
	useCase usecase.IUseCase
}

func NewCreateUserHandler(useCase usecase.IUseCase) *CreateUserHandler {
	return &CreateUserHandler{useCase: useCase}
}

func (h *CreateUserHandler) Handle(ctx context.Context, cmd commandbus.ICommand) error {
	c, ok := cmd.(CreateUserCommand)
	if !ok {
		return fmt.Errorf("invalid command type")
	}

	// Logic xử lý tạo user
	fmt.Println("Creating user:", c.Username, c.Email)
	return nil
}
