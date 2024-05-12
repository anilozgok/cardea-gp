package request

import (
	"time"
)

type NewUser struct {
	Firstname   string    `json:"firstName" validate:"required"`
	LastName    string    `json:"lastName" validate:"required"`
	DateOfBirth time.Time `json:"dateOfBirth" validate:"required"`
	Gender      string    `json:"gender" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Role        string    `json:"role" validate:"required"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ForgotPassword struct {
	Password string `json:"password" validate:"required"`
}

type VerifyPasscode struct {
	Passcode int `json:"passcode" validate:"required"`
}
