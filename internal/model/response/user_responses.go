package response

import "time"

type UserResponse struct {
	UserId      uint      `json:"userId"`
	Email       string    `json:"email"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Gender      string    `json:"gender"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Role        string    `json:"role"`
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
}

type MeResponse struct {
	UserId uint   `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type PhotosResponse struct {
	Photos []string `json:"photos"`
}
