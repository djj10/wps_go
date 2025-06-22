package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    "wps_go/proto/user"
    "wps_go/pkg/database"
    "wps_go/pkg/logic"
)

// UserServer 移至 user_logic.go
// NewUserServer 移至 user_logic.go

// 移除 CreateUser、GetUser、DeleteUser 方法

func startGRPCServer(db *gorm.DB) {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    user.RegisterUserServiceServer(s, logic.NewUserServer(db))
    log.Println("gRPC 服务器启动，监听端口: 50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}