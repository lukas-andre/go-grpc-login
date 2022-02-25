package repository

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

type AuthRepository struct {
	opts AuthRepositoryOpts
}

type AuthRepositoryOpts struct {
	Dao *gorm.DB
}

func NewAuthRepository(opts AuthRepositoryOpts) *AuthRepository {
	return &AuthRepository{opts: opts}
}

var AuthRepositorySet = wire.NewSet(wire.Struct(new(AuthRepositoryOpts), "*"), NewAuthRepository)
