package request

type CreateNewUserRequest struct {
	Firstname string `json:"firstname"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
