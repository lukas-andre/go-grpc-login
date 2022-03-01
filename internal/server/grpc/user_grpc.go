package grpc

import (
	"context"
	"fmt"
	"login_grpc/internal/models"
	"login_grpc/internal/services"
	"login_grpc/pkg"
	"time"

	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	AuthorizationHeader = "Authorization"
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
	// if _, err := s.opts.AuthService.ValidateToken(ctx); err != nil {
	// 	log.Printf("Error validating token: %v", err)
	// 	return nil, status.Error(codes.PermissionDenied, "Invalid token")
	// }

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
	user := models.User{
		Username:  in.Username,
		Password:  in.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := s.opts.UserService.CreateUser(user)

	fmt.Println(result)

	// TODO: Handle error
	if err != nil {
		return &pkg.CreateUserResponse{
			Success: false,
		}, err
	}

	return &pkg.CreateUserResponse{
		Success: true,
	}, nil
}
