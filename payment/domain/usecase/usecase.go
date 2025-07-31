package usecase

type IUseCase interface {
	AccountUsecase() IAccountUsecase
	TransactionUsecase() ITransactionUsecase
}

type UseCase struct {
	accountUsecase     IAccountUsecase
	transactionUsecase ITransactionUsecase
}

func NewUseCase() IUseCase {
	return &UseCase{
		accountUsecase:     NewAccountUsecase(),
		transactionUsecase: NewTransactionUsecase(),
	}
}

func (u *UseCase) AccountUsecase() IAccountUsecase {
	return u.accountUsecase
}

func (u *UseCase) TransactionUsecase() ITransactionUsecase {
	return u.transactionUsecase
}
