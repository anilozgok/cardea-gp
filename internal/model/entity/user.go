package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	Password    string
	Gender      string
	Role        string
	DateOfBirth time.Time
}
