package request

import "time"

type NewUserRequest struct {
	Firstname   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Gender      string    `json:"gender"`
	Height      int       `json:"height"`
	Weight      float32   `json:"weight"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
