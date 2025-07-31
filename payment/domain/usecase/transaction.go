package usecase

type ITransactionUsecase interface {
}

type TransactionUsecase struct {
}

func NewTransactionUsecase() ITransactionUsecase {
	return &TransactionUsecase{}
}
