package response

type ProfileResponse struct {
	UserId         uint     `json:"userId"`
	Bio            string   `json:"bio"`
	ProfilePicture string   `json:"profilePicture"`
	Experience     string   `json:"experience"`
	Specialization string   `json:"specializations"`
	Photos         []string `json:"photos"`
}
