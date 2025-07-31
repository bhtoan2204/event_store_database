package persistent_object

import "event_sourcing_payment/constant"

type Transaction struct {
	Base
	ID              uint   `gorm:"primaryKey"`
	TransactionCode string `gorm:"unique"`
	AccountID       uint
	Type            constant.TransactionType // Deposit, Withdraw, Transfer
	Amount          int64
	Reference       string // optional
}
