package request

type CreateDietRequest struct {
	UserId      uint   `json:"userId" validate:"required"`
	Meal        string `json:"meal" validate:"required"`
	Description string `json:"description" validate:"required"`
	Calories    int    `json:"calories" validate:"required"`
	Protein     int    `json:"protein" validate:"required"`
	Carbs       int    `json:"carbs" validate:"required"`
	Fat         int    `json:"fat" validate:"required"`
}
