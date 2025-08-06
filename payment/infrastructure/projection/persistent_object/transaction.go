package persistent_object

import "event_sourcing_payment/constant"

type Transaction struct {
	Base
	ID              int64  `gorm:"primaryKey"`
	TransactionCode string `gorm:"unique"`
	AccountNo       string
	Type            constant.TransactionType // Deposit, Withdraw, Transfer
	Amount          int64
	Reference       string // optional
}
