package dto

import "time"

type InputUser struct {
	FirstName string `json:"first_name" validate:"required,min=1"`
	LastName  string `json:"last_name" validate:"required,min=1"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
}

type OutputUser struct {
	Id string `json:"id"`
	FirstName string `json:"first_name" `
	LastName  string `json:"last_name"`
	Email     string `json:"email" `
	Password  string `json:"password" `
	Role  OutputRole `json:"role" `
	IsActive  bool `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
