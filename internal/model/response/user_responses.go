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

type PhotoResponse struct {
	PhotoId   uint      `json:"photoId"`
	PhotoURL  string    `json:"photoURL"`
	CreatedAt time.Time `json:"createdAt"`
}

type PhotosResponse struct {
	Photos []PhotoResponse `json:"photos"`
}
