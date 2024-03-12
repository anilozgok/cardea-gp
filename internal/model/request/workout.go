package request

type CreateWorkoutRequest struct {
	UserId      uint   `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Area        string `json:"area"`
	Rep         int    `json:"rep"`
	Sets        int    `json:"sets"`
}
