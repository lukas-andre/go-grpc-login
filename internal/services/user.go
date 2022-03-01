package services

import (
	"login_grpc/internal/models"
	"login_grpc/internal/repository"

	"github.com/google/wire"
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
	return s.opts.UserRepo.GetByUsername(username)
}

func (s *UserService) CreateUser(user models.User) (*models.User, error) {
	return s.opts.UserRepo.CreateUser(user)
}
