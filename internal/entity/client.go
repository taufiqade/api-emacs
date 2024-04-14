package entity

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID         uint `gorm:"primaryKey;autoIncrement"`
	ClientName string
	ClientLogo string
	Status     bool
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
