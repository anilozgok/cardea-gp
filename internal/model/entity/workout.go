package entity

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserId      uint
	CoachId     uint
	Name        string
	Exercise    uint
	Description string
	Area        string
	Rep         int
	Sets        int
}

type Exercise struct {
	gorm.Model
	BodyPart     string
	Equipment    string
	GifUrl       string
	ExerciseName string
	Target       string
}
