package jwt

type Opts struct {
	UserId uint32 `json:"userId"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}
