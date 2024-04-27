package response

type WorkoutResponse struct {
	UserId      uint   `json:"userId"`
	CoachId     uint   `json:"coachId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Area        string `json:"area"`
	Rep         int    `json:"rep"`
	Sets        int    `json:"sets"`
}
type WorkoutListResponse struct {
	Workouts []WorkoutResponse `json:"workouts"`
}
