package dao

import (
	"fmt"
	"login_grpc/internal/config"
	"login_grpc/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(config *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s", config.Connection, config.User, config.Password, config.Host, config.Port, config.Name)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
