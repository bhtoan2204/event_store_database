package dto

import (
	"errors"
	"event_sourcing_payment/constant"

	"github.com/go-playground/validator/v10"
)

type CreateTransactionRequestDto struct {
	AccountNo string  `json:"account_no" validate:"required"`
	Amount    float64 `json:"amount" validate:"required,min=1"`
	Type      string  `json:"type" validate:"required,oneof=deposit withdraw transfer"`
	Reference string  `json:"reference"`
}

func (r *CreateTransactionRequestDto) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	if r.Type == constant.TransactionTypeTransfer.String() && r.Reference == "" {
		return errors.New("reference is required when type is transfer")
	}
	return nil
}
