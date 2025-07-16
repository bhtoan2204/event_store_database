package persistent_object

import (
	"time"
)

type Base struct {
	ID        int64 `gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
