package logic

import (
	"context"
	"log"
	"wps_go/pkg/model/user"
	"wps_go/proto/user"

	"gorm.io/gorm"
)

// UserServer 实现 user.UserServiceServer 接口
type UserServer struct {
	user.UnimplementedUserServiceServer
	db *gorm.DB
}

// NewUserServer 创建新的用户服务实例
func NewUserServer(db *gorm.DB) *UserServer {
	return &UserServer{db: db}
}

// CreateUser 实现创建用户方法
func (s *UserServer) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	newUser := model_user.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	// 使用 GORM 创建用户
	result := s.db.Create(&newUser)
	if result.Error != nil {
		log.Printf("创建用户失败: %v", result.Error)
		return nil, result.Error
	}

	log.Printf("创建用户: %s, %s", req.GetName(), req.GetEmail())
	return &user.CreateUserResponse{UserId: newUser.ID.String()}, nil
}

// GetUser 实现获取用户方法
func (s *UserServer) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	// 这里实现获取用户的逻辑
	log.Printf("获取用户: %s", req.GetUserId())
	return &user.GetUserResponse{UserId: req.GetUserId(), Name: "Test User", Email: "test@example.com"}, nil
}

// DeleteUser 实现删除用户方法
func (s *UserServer) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	// 这里实现删除用户的逻辑
	log.Printf("删除用户: %s", req.GetUserId())
	return &user.DeleteUserResponse{Success: true}, nil
}
