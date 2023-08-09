package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"h_patient/config"
	"h_patient/domain/patient"
	"h_patient/pkg/dbconnect"
)

func main() {
	var err error
	config.AppConfig, err = config.LoadConfig()
	// err := godotenv.Load("../../.env")
	// err = godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	dbx, err := dbconnect.ConnectSqlx(dbconnect.DBConfig{
		Host:       config.AppConfig.Postgres.Host,
		Port:       config.AppConfig.Postgres.Port,
		Dbname:     config.AppConfig.Postgres.DbName,
		Dbuser:     config.AppConfig.Postgres.User,
		Dbpassword: config.AppConfig.Postgres.Password,
		Sslmode:    config.AppConfig.Postgres.SSLMode,
	})
	if err != nil {
		log.Fatalf("Cannot connect to DB:%v", err)
	}

	srv := grpc.NewServer()
	patient.RouterInitGrpc(srv, dbx)

	log.Println("Register RouteGRPC ...")

	listen, err := net.Listen("tcp", config.AppConfig.App.BaseGrpcPort)
	if err != nil {
		panic(err)
	}

	log.Println("running GRPC server at port", config.AppConfig.App.BaseGrpcPort)
	err = srv.Serve(listen)
	if err != nil {
		panic(err)
	}
}
