package entity

import "gorm.io/gorm"

type Diet struct {
	gorm.Model
	UserId      uint
	Meal        string
	Description string
	Calories    int
	Protein     int
	Carbs       int
	Fat         int
}
