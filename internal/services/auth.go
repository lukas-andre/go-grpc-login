package services

import (
	"login_grpc/internal/repository"

	"github.com/google/wire"
)

type AuthService struct {
	opts *AuthServiceOpts
}

type AuthServiceOpts struct {
	AuthRepo *repository.AuthRepository
}

func NewAuthService(opts *AuthServiceOpts) *AuthService {
	return &AuthService{opts: opts}
}

var AuthServiceSet = wire.NewSet(wire.Struct(new(AuthServiceOpts), "*"), NewAuthService)
