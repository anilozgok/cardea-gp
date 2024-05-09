package entity

type Profile struct {
	ID                int64    `json:"id"`
	UserID            int64    `json:"user_id"`
	Bio               string   `json:"bio"`
	ProfilePictureURL string   `json:"profile_picture_url"`
	Experience        string   `json:"experience"`
	Specialization    string   `json:"specialization"` // Use singular form
	Photos            []string `json:"photos"`         // List of photo URLs
}
