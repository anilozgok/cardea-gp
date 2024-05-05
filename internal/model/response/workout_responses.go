package response

type WorkoutResponse struct {
	UserId      uint             `json:"userId"`
	CoachId     uint             `json:"coachId"`
	Name        string           `json:"name"`
	Exercise    ExerciseResponse `json:"exercise"`
	Description string           `json:"description"`
	Area        string           `json:"area"`
	Rep         int              `json:"rep"`
	Sets        int              `json:"sets"`
}
type WorkoutListResponse struct {
	Workouts []WorkoutResponse `json:"workouts"`
}

type ExerciseResponse struct {
	BodyPart     string `json:"bodyPart"`
	Equipment    string `json:"equipment"`
	GifUrl       string `json:"gifUrl"`
	ExerciseName string `json:"exerciseName"`
	Target       string `json:"target"`
}