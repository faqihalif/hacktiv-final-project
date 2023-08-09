package patient

import (
	context "context"
	"fmt"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type handlerGrpc struct {
	svc service
}

func NewHandlerGrpc(_svc service) handlerGrpc {
	return handlerGrpc{
		svc: _svc,
	}
}

func (h handlerGrpc) Create(ctx context.Context, req *PatientProto) (emp *emptypb.Empty, err error) {
	data := createPatientRequest{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		DateOfBirth: req.DateOfBirth,
		Address:     req.Address,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
	}

	err = h.svc.CreatePatient(data)
	fmt.Printf("%+v", err)

	emp = new(emptypb.Empty)

	return
}

func (h handlerGrpc) List(ctx context.Context, emp *emptypb.Empty) (res *PatientListProto, err error) {
	data, err := h.svc.ListPatient()

	var patient []*PatientResponseProto

	for _, v := range data {
		patient = append(patient, &PatientResponseProto{
			PatientId:   v.PatientId,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			DateOfBirth: v.DateOfBirth,
			Address:     v.Address,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			UserCreate:  v.UserCreate,
		})
	}

	res = &PatientListProto{
		List: patient,
	}

	return
}

func (h handlerGrpc) Find(ctx context.Context, req *PatientFilterProto) (res *PatientListProto, err error) {
	dataRequest := filterPatientRequest{
		PatientId: req.PatientId,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	data, err := h.svc.FilterPatient(dataRequest)

	var patient []*PatientResponseProto

	for _, v := range data {
		patient = append(patient, &PatientResponseProto{
			PatientId:   v.PatientId,
			FirstName:   v.FirstName,
			LastName:    v.LastName,
			DateOfBirth: v.DateOfBirth,
			Address:     v.Address,
			Email:       v.Email,
			PhoneNumber: v.PhoneNumber,
			UserCreate:  v.UserCreate,
		})
	}

	res = &PatientListProto{
		List: patient,
	}

	return
}

func (h handlerGrpc) mustEmbedUnimplementedPatientProtoServiceServer() {
}
