package request

type MealRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Calories    float64 `json:"calories" validate:"required"`
	Protein     float64 `json:"protein"`
	Carbs       float64 `json:"carbs"`
	Fat         float64 `json:"fat"`
	Gram        float64 `json:"gram"`
}

type CreateDietRequest struct {
	UserId uint          `json:"user_id" validate:"required"`
	Name   string        `json:"name" validate:"required"`
	Meals  []MealRequest `json:"meals" validate:"required,dive"`
}
