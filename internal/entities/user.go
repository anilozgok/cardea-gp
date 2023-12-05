package entities

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Firstname   string
	LastName    string
	Email       string
	Password    string
	Height      int
	Weight      float32
	Gender      string
	Role        string
	DateOfBirth time.Time
}
