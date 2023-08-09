package patient

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) CreatePatient(data Patient) (err error) {
	query := `INSERT INTO patients (
		patient_id,
		first_name,
		last_name,
		date_of_birth,
		address,
		email,
		phone_number,
		is_active,
		create_at,
		user_create
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10
	)`

	fmt.Printf("%+v", data)

	statement, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	_, err = statement.Exec(
		data.PatientId,
		data.FirstName,
		data.LastName,
		data.DateOfBirth,
		data.Address,
		data.Email,
		data.PhoneNumber,
		data.IsActive,
		data.CreateAt,
		data.UserCreate,
	)

	return
}

func (r repository) ListPatient() (res []patientResponse, err error) {
	query := `SELECT
	patient_id,
	first_name,
	last_name,
	date_of_birth,
	address,
	email,
	phone_number,
	user_create FROM patients`

	_, err = r.db.Prepare(query)
	if err != nil {
		return res, err
	}

	err = r.db.Select(&res, query)

	return
}

func (r repository) FilterPatient(req filterPatientRequest) (res []patientResponse, err error) {
	if req.FirstName == "" && req.LastName == "" && req.Email == "" {
		query := `SELECT
			patient_id,
			first_name,
			last_name,
			date_of_birth,
			address,
			email,
			phone_number,
			user_create FROM patients WHERE patient_id=$1`

		_, err = r.db.Prepare(query)
		if err != nil {
			return res, err
		}

		err = r.db.Select(&res, query, req.PatientId)

		return
	} else {
		query := `SELECT
			patient_id,
			first_name,
			last_name,
			date_of_birth,
			address,
			email,
			phone_number,
			user_create FROM patients WHERE patient_id=$1 OR LOWER(first_name) LIKE $2 OR LOWER(last_name) LIKE $3 OR LOWER(email) LIKE $4`

		_, err = r.db.Prepare(query)
		if err != nil {
			return res, err
		}

		err = r.db.Select(&res, query, req.PatientId, "%"+req.FirstName+"%", "%"+req.LastName+"%", "%"+req.Email+"%")

		return
	}

}
