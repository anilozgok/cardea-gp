package request

type CreateProfileRequest struct { // Corrected struct name
	UserID         int64    `json:"user_id"`
	Bio            string   `json:"bio"`
	ProfilePicture string   `json:"profile_picture"`
	Experience     string   `json:"experience"`
	Specialization string   `json:"specialization"` // Use singular form
	Photos         []string `json:"photos"`         // List of photo URLs
}

type UpdateProfileRequest struct {
	Bio            string   `json:"bio"`
	ProfilePicture string   `json:"profile_picture"`
	Experience     string   `json:"experience"`
	Specialization string   `json:"specialization"` // Use singular form
	Photos         []string `json:"photos"`         // List of photo URLs
}
