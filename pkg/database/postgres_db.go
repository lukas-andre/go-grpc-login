package postgresDb

import (
	"log"
	usersModels "login_grpc/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/grpc_login"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&usersModels.User{})

	return db
}
