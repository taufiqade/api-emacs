package model

import "time"

type Client struct {
	ID         uint      `json:"id"`
	ClientName string    `json:"client_name"`
	ClientLogo string    `json:"client_logo"`
	Status     bool      `json:"status"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
