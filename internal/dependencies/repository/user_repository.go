package repoDependencies

import "login_grpc/internal/repository"

type UserRepositoryDependency func(*UserRepository)

type UserRepository struct {
	UserRepo repository.UserRepository
}

func NewUserRepositoryDepedency(deps ...UserRepositoryDependency) *UserRepository {
	d := &UserRepository{}
	for _, dep := range deps {
		dep(d)
	}
	return d
}

func WithUserRepository(userRepo repository.UserRepository) UserRepositoryDependency {
	return func(d *UserRepository) {
		d.UserRepo = userRepo
	}
}
