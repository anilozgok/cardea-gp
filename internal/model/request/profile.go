package request

type CreateProfileRequest struct {
	Bio            string  `json:"bio"`
	Height         int     `json:"height" validate:"required"`
	Weight         float32 `json:"weight" validate:"required"`
	ProfilePicture string  `json:"profilePicture"`
	Experience     string  `json:"experience" `
	Specialization string  `json:"specialization"`
	Phone          string  `json:"phone" `
	Country        string  `json:"country"`
	StateProvince  string  `json:"stateProvince"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	ZipCode        string  `json:"zipCode"`
}

type UpdateProfileRequest struct {
	Bio            string  `json:"bio"`
	Height         int     `json:"height" validate:"required"`
	Weight         float32 `json:"weight" validate:"required"`
	ProfilePicture string  `json:"profilePicture"`
	Experience     string  `json:"experience"`
	Specialization string  `json:"specialization"`
	Phone          string  `json:"phone"`
	Country        string  `json:"country"`
	StateProvince  string  `json:"stateProvince"`
	Address        string  `json:"address"`
	City           string  `json:"city"`
	ZipCode        string  `json:"zipCode"`
}

type UploadPhotoRequest struct {
	PhotoURL string `json:"photoUrl"`
}
