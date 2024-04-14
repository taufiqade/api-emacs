package model

import "time"

type GetByIdRequest struct {
	ID string `json:"id" validate:"required"`
}

type User struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        Role      `json:"role"`
	Client      Client    `json:"client"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
}

type CreateUserRequest struct {
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	RoleID      string `json:"role" validate:"required"`
	ClientID    string `json:"client" validate:"required"`
}

type UpdateUserRequest struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	Email       string `json:"email" validate:"email"`
	PhoneNumber string `json:"phone_number"`
}
