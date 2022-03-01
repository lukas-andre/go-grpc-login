package config

type DatabaseConfig struct {
	Host       string
	Port       int
	User       string
	Password   string
	Name       string
	Connection string
}

type ServerConfig struct {
	Host         string
	Port         string
	GrpcProtocol string
	GrpcPort     string
	HttpPort     string
}

type JwtConfig struct {
	Secret string
}
type config struct {
	databaseConfig *DatabaseConfig
	jwtConfig      *JwtConfig
	serverConfig   *ServerConfig
}

type Config interface {
	GetDatabaseConfig() *DatabaseConfig
	GetJWTConfig() *JwtConfig
	GetConfig() *config
	GetServerConfig() *ServerConfig
}
