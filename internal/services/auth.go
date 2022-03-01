package services

import (
	"context"
	"errors"
	"login_grpc/internal/repository"

	"github.com/google/wire"
	"golang.org/x/crypto/bcrypt"
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
	AuthServiceSet      = wire.NewSet(wire.Struct(new(AuthServiceOpts), "*"), NewAuthService)
	AuthorizationHeader = "Authorization"
)

func (s *AuthService) ValidateToken(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to get metadata")
	}

	token := md.Get(AuthorizationHeader)
	userClaims, err := s.opts.TokenHandler.ParseToken(token[len(token)-1])
	if err != nil {
		return nil, err
	}

	return userClaims, nil
}

func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *AuthService) CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}
