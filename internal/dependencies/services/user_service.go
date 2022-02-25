package serviceDependencies

import (
	"login_grpc/internal/services"
)

type userServiceDependency func(*UserServiceDependecy)

type UserServiceDependecy struct {
	UserService services.UserService
}

func NewUserServiceDependency(deps ...userServiceDependency) *UserServiceDependecy {
	d := &UserServiceDependecy{}
	for _, dep := range deps {
		dep(d)
	}
	return d
}

func WithUserService(userService services.UserService) userServiceDependency {
	return func(d *UserServiceDependecy) {
		d.UserService = userService
	}
}
