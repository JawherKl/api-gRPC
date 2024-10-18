package main

import (
    "context"
    "log"
    "net"
    "sync"

	"google.golang.org/grpc"
    "google.golang.org/grpc/codes"      // Import grpc/codes for error codes
    "google.golang.org/grpc/status"     // Import grpc/status for creating error status
    pb "api-grpc/path"
)

type server struct {
    pb.UnimplementedApiServiceServer
    users      map[string]*pb.UserResponse
    usersMutex sync.Mutex
}

// Implementation of GetUserInfo
func (s *server) GetUserInfo(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
    s.usersMutex.Lock()
    defer s.usersMutex.Unlock()

    user, exists := s.users[req.UserId]
    if !exists {
        return nil, status.Errorf(codes.NotFound, "User not found")  // Return a proper gRPC error status
    }
    return user, nil
}

// Implementation of CreateUser
func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
    s.usersMutex.Lock()
    defer s.usersMutex.Unlock()

    userId := generateUserId() // Implement this function to generate unique IDs
    user := &pb.UserResponse{
        UserId:   userId,
        UserName: req.UserName,
        UserAge:  req.UserAge,
    }
    s.users[userId] = user
    return user, nil
}

// Implementation of ListUsers
func (s *server) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
    s.usersMutex.Lock()
    defer s.usersMutex.Unlock()

    response := &pb.ListUsersResponse{}
    for _, user := range s.users {
        response.Users = append(response.Users, user)
    }
    return response, nil
}

// Generate a unique user ID (dummy implementation)
func generateUserId() string {
    // Implement your logic for generating unique user IDs
    return "unique_user_id"
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterApiServiceServer(s, &server{
        users: make(map[string]*pb.UserResponse),
    })

    log.Println("gRPC server running on port 50051...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}