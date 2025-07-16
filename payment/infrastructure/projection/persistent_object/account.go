package persistent_object

import "time"

type Account struct {
	Base
	AccountNo string `gorm:"unique"`
	UserID    uint
	Balance   int64 // value from replay event
	Currency  string
	Status    string // Active, Suspended
	CreatedAt time.Time
}

func (a *Account) TableName() string {
	return "accounts"
}
