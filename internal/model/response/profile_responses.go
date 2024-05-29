package response

import "time"

type ProfileResponse struct {
	ProfileId      uint      `json:"profileId"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Email          string    `json:"email"`
	Gender         string    `json:"gender"`
	DateOfBirth    time.Time `json:"dateOfBirth"`
	Bio            string    `json:"bio"`
	Height         int       `json:"height"`
	Weight         float32   `json:"weight"`
	ProfilePicture string    `json:"profilePicture"`
	Experience     string    `json:"experience"`
	Specialization string    `json:"specialization"`
	Phone          string    `json:"phone"`
	Country        string    `json:"country"`
	StateProvince  string    `json:"stateProvince"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	ZipCode        string    `json:"zipCode"`
}
