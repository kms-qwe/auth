package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/kms-qwe/microservices_course_auth/internal/config"
	"github.com/kms-qwe/microservices_course_auth/internal/config/env"
	"github.com/kms-qwe/microservices_course_auth/internal/model"
	"github.com/kms-qwe/microservices_course_auth/internal/storage"
	"github.com/kms-qwe/microservices_course_auth/internal/storage/postgres"
	desc "github.com/kms-qwe/microservices_course_auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	storage storage.Storage
}

// NewSever initializes a new server instance with the provided DSN for storage.
func NewSever(ctx context.Context, DSN string) (*server, error) {
	storage, err := postgres.NewPgStorage(ctx, DSN)
	if err != nil {
		return nil, err
	}
	return &server{
		storage: storage,
	}, nil
}
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	info := model.UserInfo{
		Name:     req.Info.GetName(),
		Email:    req.Info.GetEmail(),
		Password: req.Info.GetPassword(),
		Role:     int32(req.Info.GetRole()),
	}
	id, err := s.storage.AddNewUser(ctx, &info)
	if err != nil {
		log.Printf("create request error: %#v\n", err)
		return nil, err
	}
	return &desc.CreateResponse{Id: id}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.storage.GetUser(ctx, req.GetId())
	if err != nil {
		log.Printf("get request error: %#v\n", err)
		return nil, err
	}

	return &desc.GetResponse{User: &desc.User{
		Id:        user.ID,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		Info: &desc.UserInfo{
			Name:     user.Info.Name,
			Email:    user.Info.Email,
			Password: user.Info.Password,
		},
	}}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := s.storage.UpdateUserInfo(ctx, req.GetId(), req.GetName().Value, req.GetEmail().Value, int32(req.GetRole()))
	if err != nil {
		log.Printf("update request error: %#v\n", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.storage.DeleteUser(ctx, req.GetId())
	if err != nil {
		log.Printf("delete request error: %#v\n", err)
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %#v\n", err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %#v\n", err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %#v\n", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to get tcp listener: %#v\n", err)
	}

	serv, err := NewSever(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to get serv: %#v\n", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, serv)

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %#v\n", err)
	}
}
