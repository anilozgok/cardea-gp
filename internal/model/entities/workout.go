package entities

import (
	"gorm.io/gorm"
)

type Workout struct {
	gorm.Model
	UserId      uint
	Name        string
	Description string
	Area        string
	Rep         int
	Sets        int
}
