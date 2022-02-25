package services

import (
	"login_grpc/internal/models"
	"login_grpc/internal/repository"
)

type UserService struct {
	opts userServiceOpts
}

type UserServiceOpts func(*userServiceOpts)

type userServiceOpts struct {
	repo repository.UserRepository
}

func NewUserService(opts ...UserServiceOpts) UserService {
	o := &userServiceOpts{}
	for _, opt := range opts {
		opt(o)
	}
	return UserService{
		opts: *o,
	}
}

func WithUserRepository(repo repository.UserRepository) UserServiceOpts {
	return func(opt *userServiceOpts) {
		opt.repo = repo
	}
}

func (s *UserService) GetUserByUsername(username string) (models.User, error) {
	user, err := s.opts.repo.GetUserByUsernmae(username)

	return user, err
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	result, err := s.opts.repo.CreateUser(user)

	return result, err
}
