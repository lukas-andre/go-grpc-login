package config

import (
	"log"

	"github.com/spf13/viper"
)

type ViperCofig struct {
	config *config
}

func ViperConfigInit() Config {
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %v", err)
	}

	viper := viper.GetViper()

	return ViperCofig{config: loadConfig(viper)}
}

func loadConfig(v *viper.Viper) *config {
	return &config{
		databaseConfig: &DatabaseConfig{
			Host:       v.GetString("database.host"),
			Port:       v.GetInt("database.port"),
			User:       v.GetString("database.user"),
			Password:   v.GetString("database.password"),
			Name:       v.GetString("database.dbname"),
			Connection: v.GetString("database.connection"),
		},
		jwtConfig: &JwtConfig{
			Secret: v.GetString("jwt.secret"),
		},
		serverConfig: &ServerConfig{
			Host:         v.GetString("app.host"),
			GrpcProtocol: v.GetString("app.grpc_protocol"),
			GrpcPort:     v.GetString("app.grpc_port"),
			HttpPort:     v.GetString("app.http_port"),
		},
	}
}

func (v ViperCofig) GetServerConfig() *ServerConfig {
	return v.config.serverConfig
}

func (v ViperCofig) GetConfig() *config {
	return v.config
}

func (v ViperCofig) GetDatabaseConfig() *DatabaseConfig {
	return v.config.databaseConfig
}

func (v ViperCofig) GetJWTConfig() *JwtConfig {
	return v.config.jwtConfig
}
