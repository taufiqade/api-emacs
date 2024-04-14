package entity

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	DisplayName string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
