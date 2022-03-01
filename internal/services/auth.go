package services

import (
	"context"
	"errors"
	"login_grpc/internal/repository"

	"github.com/google/wire"
	"google.golang.org/grpc/metadata"
)

type AuthService struct {
	opts *AuthServiceOpts
}

type AuthServiceOpts struct {
	TokenHandler *TokenHandler
	AuthRepo     *repository.AuthRepository
}

func NewAuthService(opts *AuthServiceOpts) *AuthService {
	return &AuthService{opts: opts}
}

var (
	AuthorizationHeader = "Authorization"
)

var AuthServiceSet = wire.NewSet(wire.Struct(new(AuthServiceOpts), "*"), NewAuthService)

func (s *AuthService) ValidateToken(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to get metadata")
	}

	token := md.Get(AuthorizationHeader)
	userClaims, err := s.opts.TokenHandler.ParseToken(token[0])

	if err != nil {
		return nil, err
	}

	// s.opts.AuthRepo.UpdateLastLogin(userClaims.UserID)

	return userClaims, nil
}
