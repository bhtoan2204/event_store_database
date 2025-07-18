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

func (b *Base) ID() int64 {
	return b.id
}

func (b *Base) CreatedAt() time.Time {
	return b.createdAt
}

func (b *Base) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b *Base) DeletedAt() *time.Time {
	return b.deletedAt
}

func (b *Base) SetID(id int64) {
	b.id = id
}

func (b *Base) SetCreatedAt(createdAt time.Time) {
	b.createdAt = createdAt
}

func (b *Base) SetUpdatedAt(updatedAt time.Time) {
	b.updatedAt = updatedAt
}
