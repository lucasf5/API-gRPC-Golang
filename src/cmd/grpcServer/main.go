package main

import (
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpcInternal "grpcTreino/src/grpc"
	"grpcTreino/src/service"
)

func main() {
	categoryDB := service.NewCategoryService()
	grpcServer := grpc.NewServer()
	grpcInternal.RegisterCategoryServiceServer(grpcServer, categoryDB)
	reflection.Register(grpcServer)
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(listen)

}
