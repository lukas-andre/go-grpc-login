package repository

import (
	"login_grpc/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	opts userOpts
}

type UserRepositoryOpts func(*userOpts)

type userOpts struct {
	Dao *gorm.DB
}

func NewUserRepository(opts ...UserRepositoryOpts) UserRepository {
	o := &userOpts{}
	for _, opt := range opts {
		opt(o)
	}
	return UserRepository{
		opts: *o,
	}
}

func WithUserRepositoryDao(dao *gorm.DB) UserRepositoryOpts {
	return func(opt *userOpts) {
		opt.Dao = dao
	}
}

func (repo *UserRepository) GetUserByUsernmae(username string) (models.User, error) {
	user := models.User{}

	tx := repo.opts.Dao.Find(&user, "username = ?", username)

	if tx.Error != nil {
		return user, tx.Error
	}

	return user, nil
}

func (repo *UserRepository) CreateUser(user models.User) (models.User, error) {
	tx := repo.opts.Dao.Create(&user)

	if tx.Error != nil {
		return user, tx.Error
	}

	return user, nil
}
