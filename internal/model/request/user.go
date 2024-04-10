package request

import "time"

type NewUserRequest struct {
	Firstname   string    `json:"firstName" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	Gender      string    `json:"gender" validate:"required"`
	Height      int       `json:"height" validate:"required"`
	Weight      float32   `json:"weight" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Role        string    `json:"role" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
