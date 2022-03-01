package grpc

import (
	"context"
	"login_grpc/internal/models"
	"login_grpc/internal/services"
	"login_grpc/pkg"
	"time"

	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	pkg.UnimplementedUserServiceServer
	opts *UserServerOpts
}

type UserServerOpts struct {
	UserService *services.UserService
	AuthService *services.AuthService
}

func NewUserServiceServer(opts *UserServerOpts) *UserServiceServer {
	return &UserServiceServer{
		opts: opts,
	}
}

var UserServiceServerSet = wire.NewSet(wire.Struct(new(UserServerOpts), "*"), NewUserServiceServer)

func (s *UserServiceServer) GetUserInfo(ctx context.Context, in *pkg.GetUserInfoRequest) (*pkg.GetUserInfoResponse, error) {
	user, _ := s.opts.UserService.GetUserByUsername(in.Username)
	if user.ID == 0 {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pkg.GetUserInfoResponse{
		Username: user.Username,
		Id:       int32(user.ID),
		CreateAt: user.CreatedAt.Local().String(),
	}, nil
}

func (s *UserServiceServer) CreateUser(ctx context.Context, in *pkg.CreateUserRequest) (*pkg.CreateUserResponse, error) {
	passwordHash, err := s.opts.AuthService.HashPassword(in.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to hash password")
	}

	user := models.User{
		Username:  in.Username,
		Password:  passwordHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err = s.opts.UserService.CreateUser(user); err != nil {
		return &pkg.CreateUserResponse{
			Success: false,
		}, status.Error(codes.Internal, err.Error())
	}

	return &pkg.CreateUserResponse{
		Success: true,
	}, nil
}
