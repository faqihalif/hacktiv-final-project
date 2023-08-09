package patient

type service struct {
	repo iRepo
}

type iRepo interface {
	CreatePatient(data Patient) (err error)
	ListPatient() (res []patientResponse, err error)
	FilterPatient(req filterPatientRequest) (res []patientResponse, err error)
}

func NewService(repo iRepo) service {
	return service{
		repo: repo,
	}
}

func (s service) CreatePatient(req createPatientRequest) (err error) {
	err = s.repo.CreatePatient(req.ConvertIntoPatient())

	return
}

func (s service) ListPatient() (res []patientResponse, err error) {
	res, err = s.repo.ListPatient()

	return
}

func (s service) FilterPatient(req filterPatientRequest) (res []patientResponse, err error) {
	res, err = s.repo.FilterPatient(req)

	return
}
