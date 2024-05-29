package entity

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserId            uint
	Bio               string
	Height            int
	Weight            float32
	ProfilePictureURL string
	Experience        string
	Specialization    string
	Phone             string
	Country           string
	StateProvince     string
	Address           string
	City              string
	ZipCode           string
}

type Image struct {
	gorm.Model
	UserId    uint
	ImageName string
	ImagePath string
}
