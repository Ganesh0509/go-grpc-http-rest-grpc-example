package services

import (
	"context"
	"log"

	pb "github.com/Ganesh0509/go-grpc-http-rest-grpc-example/pkg/api/v1"
)

type server struct{}

//NewServer ...
func NewServer() *server {
	return &server{}
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("request: ", in.Name)
	return &pb.HelloReply{Message: "hello, " + in.Name}, nil
}
