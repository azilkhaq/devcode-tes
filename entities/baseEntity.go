package entities

import "time"

type Base struct {
	CreatedAt   *time.Time `gorm:"default:current_timestamp" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"default:current_timestamp" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"default:current_timestamp" json:"deleted_at"`
}