package services

import (
	"login_grpc/internal/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenHandler struct {
	signingKey string
}

func NewTokenHandler(signingKey string) *TokenHandler {
	return &TokenHandler{signingKey: signingKey}
}

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

	return token.SignedString(t.signingKey)
}
