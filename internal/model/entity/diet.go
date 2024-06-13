package entity

import (
	"gorm.io/gorm"
	"time"
)

type Diet struct {
	gorm.Model
	UserId  uint   `json:"user_id"`
	CoachId uint   `json:"coach_id"`
	Name    string `json:"name"`
	Meals   []Meal `gorm:"foreignKey:DietID" json:"meals"`
}

type Meal struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	DietID      uint           `json:"diet_id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Calories    float64        `json:"calories"`
	Protein     float64        `json:"protein"`
	Carbs       float64        `json:"carbs"`
	Fat         float64        `json:"fat"`
	Gram        float64        `json:"gram"`
}
