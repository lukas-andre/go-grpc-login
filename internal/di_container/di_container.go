//go:build wireinject
// +build wireinject

package di

import (
	"login_grpc/internal/config"
	"login_grpc/internal/repository"
	"login_grpc/internal/server/grpc"
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

func InitializeUserServiceServer(us *services.UserService, as *services.AuthService) *grpc.UserServiceServer {
	wire.Build(
		grpc.UserServiceServerSet,
	)

	return &grpc.UserServiceServer{}
}

// Auth Dependencies Injection
func InitializeAuthRepository(dao *gorm.DB) *repository.AuthRepository {
	wire.Build(
		repository.AuthRepositorySet,
	)

	return &repository.AuthRepository{}
}

func InitializeAuthService(ar *repository.AuthRepository, th *services.TokenHandler) *services.AuthService {
	wire.Build(
		services.AuthServiceSet,
	)

	return &services.AuthService{}
}

func InitializeAuthServiceServer(as *services.AuthService, us *services.UserService, th *services.TokenHandler) *grpc.AuthServiceServer {
	wire.Build(
		grpc.AuthServiceServerSet,
	)

	return &grpc.AuthServiceServer{}
}

// Token Dependencies Injection
func InitializeTokenHandler(secret *config.JwtConfig) *services.TokenHandler {
	wire.Build(services.TokenHandlerSet)

	return &services.TokenHandler{}
}
