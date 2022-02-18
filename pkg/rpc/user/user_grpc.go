package userGrpc

import (
	"context"
	"fmt"
	usersModels "login_grpc/pkg/models"
	"login_grpc/pkg/protogen/loginpb"
	"time"

	"gorm.io/gorm"
)

type UserServiceServer struct {
	loginpb.UnimplementedUserServiceServer
	dependencies
}

type Dependencies func(*dependencies)

type dependencies struct {
	dbClient *gorm.DB
}

func NewUserService(deps ...Dependencies) *UserServiceServer {
	d := &dependencies{}
	for _, dep := range deps {
		dep(d)
	}
	return &UserServiceServer{
		dependencies: *d,
	}
}

func WithDbClient(dbClient *gorm.DB) Dependencies {
	return func(d *dependencies) {
		d.dbClient = dbClient
	}
}

func (s *UserServiceServer) GetUserInfo(ctx context.Context, in *loginpb.GetUserInfoRequest) (*loginpb.GetUserInfoResponse, error) {
	// TODO: Add Auth Validation
	user := usersModels.User{}

	s.dbClient.Find(&user, "username = ?", in.Username)

	// TODO: Invoke gRPC error
	if user.ID == 0 {
		return nil, nil
	}

	return &loginpb.GetUserInfoResponse{
		Username: user.Username,
		Id:       int32(user.ID),
		CreateAt: user.CreatedAt.Local().String(),
	}, nil
}

func (s *UserServiceServer) CreateUser(ctx context.Context, in *loginpb.CreateUserRequest) (*loginpb.CreateUserResponse, error) {
	user := usersModels.User{
		Username:  in.Username,
		Password:  in.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// TODO: Handle error
	result := s.dbClient.Create(&user)

	fmt.Println(result)

	if result.Error != nil {
		return &loginpb.CreateUserResponse{
			Success: false,
		}, result.Error
	}

	return &loginpb.CreateUserResponse{
		Success: true,
	}, nil
}
