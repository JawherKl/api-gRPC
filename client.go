package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "api-grpc/path"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewApiServiceClient(conn)

    // Create a new user
    createUser(client, "Alice", 25)

    // Get user information
    getUser(client, "unique_user_id") // Replace with actual user ID

    // List all users
    listUsers(client)
}

// Function to create a user
func createUser(client pb.ApiServiceClient, name string, age int32) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := client.CreateUser(ctx, &pb.CreateUserRequest{UserName: name, UserAge: age})
    if err != nil {
        log.Fatalf("Error calling CreateUser: %v", err)
    }
    log.Printf("User created: ID=%s, Name=%s, Age=%d", res.UserId, res.UserName, res.UserAge)
}

// Function to get user info
func getUser(client pb.ApiServiceClient, userID string) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := client.GetUserInfo(ctx, &pb.UserRequest{UserId: userID})
    if err != nil {
        log.Fatalf("Error calling GetUserInfo: %v", err)
    }
    log.Printf("User info: ID=%s, Name=%s, Age=%d", res.UserId, res.UserName, res.UserAge)
}

// Function to list all users
func listUsers(client pb.ApiServiceClient) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    res, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
    if err != nil {
        log.Fatalf("Error calling ListUsers: %v", err)
    }

    log.Println("List of users:")
    for _, user := range res.Users {
        log.Printf("User ID=%s, Name=%s, Age=%d", user.UserId, user.UserName, user.UserAge)
    }
}