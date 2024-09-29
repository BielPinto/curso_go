package main

import (
	"database/sql"
	"net"

	"github.com/BielPinto/curso_go/12-grpc/internal/database"
	"github.com/BielPinto/curso_go/12-grpc/internal/pb"
	"github.com/BielPinto/curso_go/12-grpc/internal/services"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	categoryDb := database.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
