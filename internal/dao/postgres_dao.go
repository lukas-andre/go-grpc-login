package dao

import (
	"fmt"
	"login_grpc/internal/models"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBConfig struct {
	Connection string
	Host       string
	Port       int
	Username   string
	Password   string
	Dbname     string
}

func UnmarshalPostgreDBConfig(config *viper.Viper) *PostgresDBConfig {
	return &PostgresDBConfig{
		Connection: config.GetString("database.connection"),
		Host:       config.GetString("database.host"),
		Port:       config.GetInt("database.port"),
		Username:   config.GetString("database.username"),
		Password:   config.GetString("database.password"),
		Dbname:     config.GetString("database.dbname"),
	}
}

func NewPostgresDB(config *viper.Viper) (*gorm.DB, error) {
	c := UnmarshalPostgreDBConfig(config)

	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s", c.Connection, c.Username, c.Password, c.Host, c.Port, c.Dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
