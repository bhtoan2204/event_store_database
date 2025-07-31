package usecase

type IAccountUsecase interface {
}

type AccountUsecase struct {
}

func NewAccountUsecase() IAccountUsecase {
	return &AccountUsecase{}
}
