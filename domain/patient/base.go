package patient

import (
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func RouterInitGrpc(s *grpc.Server, db *sqlx.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	handler := NewHandlerGrpc(svc)

	RegisterPatientProtoServiceServer(s, handler)
}
