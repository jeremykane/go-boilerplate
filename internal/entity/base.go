package entity

import "time"

type BaseEntity struct {
	ID        int64     `gorm:"primaryKey, column:id" json:"id"`
	Uuid      string    `gorm:"column:uuid" json:"uuid"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}
