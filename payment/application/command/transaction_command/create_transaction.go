package transaction_command

import (
	"context"
	"event_sourcing_payment/domain/usecase"
	"event_sourcing_payment/dto"
	"event_sourcing_payment/package/commandbus"
	"event_sourcing_payment/package/logger"
	"fmt"

	"go.uber.org/zap"
)

type CreateTransactionCommand struct {
	AccountNo string
	Amount    float64
	Type      string
	Reference string
}

func (c CreateTransactionCommand) CommandName() string {
	return "CreateTransactionCommand"
}

type CreateTransactionHandler struct {
	useCase usecase.ITransactionUsecase
}

func NewCreateTransactionHandler(useCase usecase.ITransactionUsecase) *CreateTransactionHandler {
	return &CreateTransactionHandler{useCase: useCase}
}

func (h CreateTransactionHandler) Handle(ctx context.Context, cmd commandbus.ICommand) error {
	log := logger.FromContext(ctx)
	c, ok := cmd.(CreateTransactionCommand)
	if !ok {
		log.Error("Invalid command type", zap.Any("command", cmd))
		return fmt.Errorf("invalid command type")
	}
	return h.useCase.CreateTransaction(ctx,
		&dto.CreateTransactionRequestDto{
			AccountNo: c.AccountNo,
			Amount:    c.Amount,
			Type:      c.Type,
			Reference: c.Reference,
		})
}
