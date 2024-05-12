package request

type CreateProfileRequest struct { // Corrected struct name
	UserId         uint     `json:"userId"`
	Bio            string   `json:"bio"`
	ProfilePicture string   `json:"profilePicture"`
	Experience     string   `json:"experience"`
	Specialization string   `json:"specialization"` // Use singular form
	Photos         []string `json:"photos"`         // List of photo URLs
}

type UpdateProfileRequest struct {
	Bio            string   `json:"bio"`
	ProfilePicture string   `json:"profilePicture"`
	Experience     string   `json:"experience"`
	Specialization string   `json:"specialization"` // Use singular form
	Photos         []string `json:"photos"`         // List of photo URLs
}
