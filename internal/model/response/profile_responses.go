package response

type ProfileResponse struct {
	ID             int64    `json:"id"`
	UserID         int64    `json:"user_id"`
	Bio            string   `json:"bio"`
	ProfilePicture string   `json:"profile_picture"`
	Experience     string   `json:"experience"`
	Specialization string   `json:"specializations"`
	Photos         []string `json:"photos"`
}
