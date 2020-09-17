package main

import (
	"log"
	"net"

	pb "github.com/Ganesh0509/go-grpc-http-rest-grpc-example/pkg/api/v1"
	"github.com/Ganesh0509/go-grpc-http-rest-grpc-example/services"

	"google.golang.org/grpc"
)

const (
	PORT = ":9192"
)

func main() {
	listeener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Got some issue")
	}

	serve := grpc.NewServer()

	pb.RegisterGreeterServer(serve, services.NewServer())
	log.Println("rpc services started, listen on localhost:9192")
	serve.Serve(listeener)
}
