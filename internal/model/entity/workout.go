package entity

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserId      uint
	CoachId     uint
	Name        string
	ExerciseId  uint
	Description string
	Area        string
	Rep         int
	Sets        int
}

type Exercise struct {
	gorm.Model
	Name      string
	BodyPart  string
	Target    string
	Equipment string
	Gif       string
}
