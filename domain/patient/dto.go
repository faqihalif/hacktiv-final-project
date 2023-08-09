package patient

import (
	"encoding/hex"
	"time"

	"github.com/google/uuid"
)

type createPatientRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	UserCreate  string `json:"user_create"`
}

func (c createPatientRequest) ConvertIntoPatient() Patient {
	id := uuid.New()
	idConvert := hex.EncodeToString(id[:8])

	return Patient{
		PatientId:   idConvert,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		DateOfBirth: c.DateOfBirth,
		Address:     c.Address,
		Email:       c.Email,
		PhoneNumber: c.PhoneNumber,
		IsActive:    true,
		CreateAt:    time.Now(),
		UserCreate:  c.UserCreate,
	}
}

type patientResponse struct {
	PatientId   string `db:"patient_id" json:"patient_id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	Address     string `db:"address" json:"address"`
	Email       string `db:"email" json:"email"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
	UserCreate  string `db:"user_create" json:"user_create"`
}

type filterPatientRequest struct {
	PatientId string `db:"patient_id" json:"patient_id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
}
