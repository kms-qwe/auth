package main

import (
	"context"
	"log"
	"time"

	desc "github.com/kms-qwe/microservices_course_auth/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:9001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("failed to close conntection: %v", err)
		}
	}()

	client := desc.NewUserV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Get(ctx, &desc.GetRequest{Id: 0})
	if err != nil {
		log.Fatalf("failed to get user by id: %v", err)
	}

	log.Printf("user info: %#v\n", response)
}
