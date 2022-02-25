//go:build wireinject
// +build wireinject

package di

import (
	"login_grpc/internal/app/grpcs"
	"login_grpc/internal/repository"
	"login_grpc/internal/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// User Dependencies Injection
func InitializeUserRepository(dao *gorm.DB) *repository.UserRepository {
	wire.Build(
		repository.UserRepositorySet,
	)

	return &repository.UserRepository{}
}

func InitializeUserService(ur *repository.UserRepository) *services.UserService {
	wire.Build(
		services.UserServiceSet,
	)

	return &services.UserService{}
}

func InitializeUserServiceServer(us *services.UserService) *grpcs.UserServiceServer {
	wire.Build(
		grpcs.UserServiceServerSet,
	)

	return &grpcs.UserServiceServer{}
}

// Auth Dependencies Injection
func InitializeAuthRepository(dao *gorm.DB) *repository.AuthRepository {
	wire.Build(
		repository.AuthRepositorySet,
	)

	return &repository.AuthRepository{}
}

func InitializeAuthService(ar *repository.AuthRepository) *services.AuthService {
	wire.Build(
		services.AuthServiceSet,
	)

	return &services.AuthService{}
}

func InitializeAuthServiceServer(as *services.AuthService, us *services.UserService, th *services.TokenHandler) *grpcs.AuthServiceServer {
	wire.Build(
		grpcs.AuthServiceServerSet,
	)

	return &grpcs.AuthServiceServer{}
}

// Token Dependencies Injection
func InitializeTokenHandler(secret string) *services.TokenHandler {
	wire.Build(services.TokenHandlerSet)

	return &services.TokenHandler{}
}
