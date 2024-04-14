package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Password    string
	Name        string
	Email       string
	PhoneNumber string
	RoleID      string
	ClientID    string
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	Role   Role   `gorm:"foreignKey:role_id;references:id"`
	Client Client `gorm:"foreignKey:client_id;references:id"`
}
