package request

import "github.com/anilozgok/cardea-gp/internal/model/entity"

type UpdateDietRequest struct {
	ID    uint          `json:"id" validate:"required"`
	Name  string        `json:"name" validate:"required"`
	Meals []entity.Meal `json:"meals" validate:"required,dive"`
}
