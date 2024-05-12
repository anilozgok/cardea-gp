package response

type ProfileResponse struct {
	UserId         uint    `json:"userId"`
	Bio            string  `json:"bio"`
	Height         int     `json:"height"`
	Weight         float32 `json:"weight"`
	ProfilePicture string  `json:"profilePicture"`
	Experience     string  `json:"experience"`
	Specialization string  `json:"specializations"`
}
