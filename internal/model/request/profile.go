package request

type CreateProfileRequest struct {
	Bio            string  `json:"bio" validate:"required"`
	Height         int     `json:"height" validate:"required"`
	Weight         float32 `json:"weight" validate:"required"`
	ProfilePicture string  `json:"profilePicture" validate:"required"`
	Experience     string  `json:"experience" validate:"required"`
	Specialization string  `json:"specialization" validate:"required"`
}

type UpdateProfileRequest struct {
	Bio            string  `json:"bio" validate:"required"`
	Height         int     `json:"height" validate:"required"`
	Weight         float32 `json:"weight" validate:"required"`
	ProfilePicture string  `json:"profilePicture" validate:"required"`
	Experience     string  `json:"experience" validate:"required"`
	Specialization string  `json:"specialization" validate:"required"`
}

type UploadPhotoRequest struct {
	PhotoURL string `json:"photoUrl"`
}
