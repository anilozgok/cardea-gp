package response

type WorkoutResponse struct {
	WorkoutId   uint   `json:"workoutId"`
	UserId      uint   `json:"userId"`
	CoachId     uint   `json:"coachId"`
	Name        string `json:"name"`
	Exercise    uint   `json:"exercise"`
	Description string `json:"description"`
	Area        string `json:"area"`
	Rep         int    `json:"rep"`
	Sets        int    `json:"sets"`
}
type WorkoutListResponse struct {
	Workouts []WorkoutResponse `json:"workouts"`
}

type ExerciseListResponse struct {
	Exercises []ExerciseResponse `json:"exercises"`
}

type ExerciseResponse struct {
	ExerciseId uint   `json:"exerciseId"`
	BodyPart   string `json:"bodyPart"`
	Equipment  string `json:"equipment"`
	Gif        string `json:"gifUrl"`
	Name       string `json:"exerciseName"`
	Target     string `json:"target"`
}
