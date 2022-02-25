package dao

import (
	"login_grpc/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDAO struct {
	db *gorm.DB
}

func NewPostgresDAO(db *gorm.DB) *PostgresDAO {
	return &PostgresDAO{db: db}
}

func NewPostgresDB() (*gorm.DB, error) {
	dsn := "postgres://postgres:postgres@localhost:5432/grpc_login"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
