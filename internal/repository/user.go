package repository

import (
	"login_grpc/internal/models"

	"github.com/google/wire"
	"gorm.io/gorm"
)

type UserRepository struct {
	opts *UserRepositoryOpts
}

type UserRepositoryOpts struct {
	Dao *gorm.DB
}

func NewUserRepository(opts *UserRepositoryOpts) *UserRepository {
	return &UserRepository{opts: opts}
}

var UserRepositorySet = wire.NewSet(wire.Struct(new(UserRepositoryOpts), "*"), NewUserRepository)

func (repo *UserRepository) GetByUsername(username string) (*models.User, error) {
	user := models.User{}

	tx := repo.opts.Dao.Find(&user, "username = ?", username)

	if tx.Error != nil {
		return &user, tx.Error
	}

	return &user, nil
}

func (repo *UserRepository) CreateUser(user models.User) (*models.User, error) {
	tx := repo.opts.Dao.Create(&user)

	if tx.Error != nil {
		return &user, tx.Error
	}

	return &user, nil
}
