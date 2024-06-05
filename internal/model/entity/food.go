package entity

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name           string  `json:"name"`
	AvgServingSize float64 `json:"avg_serving_size"`
	Calories       float64 `json:"calories"`
	Category       string  `json:"category"`
}
