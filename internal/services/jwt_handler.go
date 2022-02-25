package services

import (
	"login_grpc/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
)

type TokenHandler struct {
	opts TokenHandlerOpts
}

type TokenHandlerOpts struct {
	SigningKey string
}

func NewTokenHandler(opts TokenHandlerOpts) *TokenHandler {
	return &TokenHandler{opts: opts}
}

var TokenHandlerSet = wire.NewSet(wire.Struct(new(TokenHandlerOpts), "*"), NewTokenHandler)

type UserInfo struct {
	Username string
	Id       int
}

type UserClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func (t *TokenHandler) CreateToken(user models.User) (string, error) {
	claims := &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "go-grpc",
		},
		UserInfo{user.Username, user.ID},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(t.opts.SigningKey)
}
