package main

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"exercise1/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
	mu     sync.Mutex
	users  map[int64]*pb.User
	nextID int64
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nextID++
	user := &pb.User{Id: s.nextID, Name: req.Name, Email: req.Email}
	s.users[s.nextID] = user
	return user, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, ok := s.users[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user %d not found", req.Id)
	}
	return user, nil
}

func (s *server) ListUsers(req *pb.ListUsersRequest, stream pb.UserService_ListUsersServer) error {
	s.mu.Lock()
	users := make([]*pb.User, 0, len(s.users))
	for _, u := range s.users {
		users = append(users, u)
	}
	s.mu.Unlock()

	for _, u := range users {
		time.Sleep(100 * time.Millisecond)
		if err := stream.Send(u); err != nil {
			return err
		}
	}
	return nil
}

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("method=%s duration=%s err=%v", info.FullMethod, time.Since(start), err)
	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
	pb.RegisterUserServiceServer(s, &server{
		users: make(map[int64]*pb.User),
	})

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
