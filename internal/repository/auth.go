package repository

import (
	"gorm.io/gorm"
)

type AuthRepository struct {
	opts authOpts
}

type AuthRepositoryOpts func(*authOpts)

type authOpts struct {
	Dao *gorm.DB
}

func NewAuthRepository(deps ...AuthRepositoryOpts) *AuthRepository {
	d := &authOpts{}
	for _, dep := range deps {
		dep(d)
	}
	return &AuthRepository{
		opts: *d,
	}
}

func WithAuthRepositoryDao(dao *gorm.DB) AuthRepositoryOpts {
	return func(opt *authOpts) {
		opt.Dao = dao
	}
}
