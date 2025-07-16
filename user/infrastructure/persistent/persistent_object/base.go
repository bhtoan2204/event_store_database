package persistent_object

import (
	"time"
)

type Base struct {
	id        int64      `gorm:"primaryKey;autoIncrement"`
	createdAt time.Time  `gorm:"autoCreateTime"`
	updatedAt time.Time  `gorm:"autoUpdateTime"`
	deletedAt *time.Time `gorm:"index"`
}
