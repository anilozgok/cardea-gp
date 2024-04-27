package entity

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserId      uint
	CoachId     uint
	Name        string
	Description string
	Area        string
	Rep         int
	Sets        int
}
