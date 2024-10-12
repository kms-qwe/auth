package main

import (
	"context"
	"log"
	"net"

	desc "github.com/kms-qwe/microservices_course_auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	address  = "localhost"
	grpcPort = "9001"
)

type server struct {
	desc.UnimplementedUserV1Server
}

func (s *server) Create(_ context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("get create request: %#v:\n", req)
	return &desc.CreateResponse{Id: 0}, nil
}
func (s *server) Get(_ context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("get gete request: %#v:", req)
	return &desc.GetResponse{User: &desc.User{}}, nil
}
func (s *server) Update(_ context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("get update request: %#v\n", req)
	return &emptypb.Empty{}, nil
}
func (s *server) Delete(_ context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("get delete request: %#v\n", req)
	return &emptypb.Empty{}, nil
}
func main() {
	lis, err := net.Listen("tcp", address+":"+grpcPort)
	if err != nil {
		log.Fatalf("falied to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
