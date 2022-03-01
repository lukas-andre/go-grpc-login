package services

import (
	"fmt"
	"login_grpc/internal/models"
	"login_grpc/internal/repository"

	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	opts *UserServiceOpts
}

type UserServiceOpts struct {
	UserRepo *repository.UserRepository
}

func NewUserService(opts *UserServiceOpts) *UserService {
	return &UserService{opts: opts}
}

var UserServiceSet = wire.NewSet(wire.Struct(new(UserServiceOpts), "*"), NewUserService)

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.opts.UserRepo.GetByUsername(username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to get user: %s", username))
	}
	return user, nil
}

func (s *UserService) CreateUser(u models.User) (*models.User, error) {
	user, err := s.opts.UserRepo.GetByUsername(u.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Failed to create user: %s", u.Username))
	}

	if user.ID != 0 {
		return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("User %s already exists", u.Username))
	}

	createdUser, err := s.opts.UserRepo.CreateUser(u)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create user")
	}

	return createdUser, nil
}
