# gRPC API Project

This project demonstrates the use of gRPC (Remote Procedure Call) for building a simple API that handles user information. The API provides functionality for creating users, retrieving user information, and listing all users.

## Prerequisites

Before running the project, ensure that you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.18 or later)
- [Protocol Buffers Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
- gRPC Go plugin for protoc:  
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

## Project Structure

```
├── api/
│   └── api.proto           # Protobuf file defining gRPC API
├── generated/
│   └── api.pb.go           # Generated Go code from api.proto
├── main.go                 # Main server implementation
├── README.md               # This file
└── go.mod                  # Go module file
```

## Protobuf Definition

The gRPC service and message types are defined in the `api/api.proto` file. Here is an example of the protobuf definition:

```proto
syntax = "proto3";

package api;

service ApiService {
  rpc GetUserInfo (UserRequest) returns (UserResponse);
  rpc CreateUser (CreateUserRequest) returns (UserResponse);
  rpc ListUsers (ListUsersRequest) returns (ListUsersResponse);
}

message UserRequest {
  string user_id = 1;
}

message CreateUserRequest {
  string user_name = 1;
  int32 user_age = 2;
}

message UserResponse {
  string user_id = 1;
  string user_name = 2;
  int32 user_age = 3;
}

message ListUsersRequest {}

message ListUsersResponse {
  repeated UserResponse users = 1;
}
```

## Running the Project

1. **Clone the repository**:
    ```bash
    git clone https://github.com/JawherKl/api-gRPC
    cd api-gRPC
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Generate gRPC code from the `.proto` file** (if not done yet):
    ```bash
    protoc --go_out=generated --go-grpc_out=generated api/api.proto
    ```

4. **Run the gRPC server**:
    ```bash
    go run main.go
    ```

    The gRPC server will start listening on `localhost:50051`.

## API Endpoints

### 1. **Get User Information**

- **Method**: `GetUserInfo`
- **Request**:
    ```json
    {
      "user_id": "string"
    }
    ```
- **Response**:
    ```json
    {
      "user_id": "string",
      "user_name": "string",
      "user_age": "int"
    }
    ```

### 2. **Create User**

- **Method**: `CreateUser`
- **Request**:
    ```json
    {
      "user_name": "string",
      "user_age": "int"
    }
    ```
- **Response**:
    ```json
    {
      "user_id": "string",
      "user_name": "string",
      "user_age": "int"
    }
    ```

### 3. **List All Users**

- **Method**: `ListUsers`
- **Request**: None
- **Response**:
    ```json
    {
      "users": [
        {
          "user_id": "string",
          "user_name": "string",
          "user_age": "int"
        }
      ]
    }
    ```

## Testing with gRPC Client

You can test the gRPC endpoints using a gRPC client like [Evans](https://github.com/ktr0731/evans) or [BloomRPC](https://github.com/uw-labs/bloomrpc).

### Example using Evans:

1. Start the Evans CLI:
    ```bash
    evans --host localhost --port 50051 -r
    ```

2. Call the `CreateUser` endpoint:
    ```
    call CreateUser
    {
      "user_name": "John Doe",
      "user_age": 30
    }
    ```

3. Retrieve the user's information:
    ```
    call GetUserInfo
    {
      "user_id": "generated-user-id"
    }
    ```

4. List all users:
    ```
    call ListUsers
    ```

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

### Explanation:
- **Project Structure**: Describes the key components of the project.
- **Running the Project**: Steps to run the server.
- **API Endpoints**: Provides details on how to call the gRPC endpoints.
- **Testing with gRPC Client**: Explains how to test the API using Evans or BloomRPC.

You can customize the file further as needed. Let me know if you'd like additional changes!
