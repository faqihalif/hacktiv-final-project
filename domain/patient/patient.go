package patient

import (
	"time"
)

type Patient struct {
	PatientId   string
	FirstName   string
	LastName    string
	DateOfBirth string
	Address     string
	Email       string
	PhoneNumber string
	IsActive    bool
	CreateAt    time.Time
	UserCreate  string
}
