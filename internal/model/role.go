package model

import "time"

type Role struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
