package main

import (
	"context"
	"io"
	"log"

	"exercise1/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx := context.Background()

	created, err := client.CreateUser(ctx, &pb.CreateUserRequest{
		Name:  "Viktor",
		Email: "viktor@test.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created: id=%d name=%s email=%s", created.Id, created.Name, created.Email)

	user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: created.Id})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got:     id=%d name=%s email=%s", user.Id, user.Name, user.Email)

	_, err = client.GetUser(ctx, &pb.GetUserRequest{Id: 999})
	if err != nil {
		log.Printf("Expected error: %v", err)
	}

	stream, err := client.ListUsers(ctx, &pb.ListUsersRequest{PageSize: 10})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Streaming users:")
	for {
		u, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("  stream: id=%d name=%s email=%s", u.Id, u.Name, u.Email)
	}
}
