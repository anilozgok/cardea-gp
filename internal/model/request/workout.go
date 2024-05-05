package request

type CreateWorkout struct {
	UserId      uint   `json:"userId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Exercise    uint   `json:"exercise" validate:"required"`
	Description string `json:"description" validate:"required"`
	Area        string `json:"area" validate:"required"`
	Rep         int    `json:"rep" validate:"required"`
	Sets        int    `json:"sets" validate:"required"`
}
