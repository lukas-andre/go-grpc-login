package services

import (
	"fmt"
	"login_grpc/internal/config"
	"login_grpc/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
)

type TokenHandler struct {
	opts TokenHandlerOpts
}

type TokenHandlerOpts struct {
	JwtConfig *config.JwtConfig
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

func (t *TokenHandler) CreateToken(user *models.User) (string, error) {
	claims := &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "go-grpc",
		},
		UserInfo{user.Username, user.ID},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.opts.JwtConfig.Secret))
}

func (t *TokenHandler) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(t.opts.JwtConfig.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
