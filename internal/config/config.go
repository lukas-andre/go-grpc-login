package config

type ServerConfig struct {
	Host         string
	Port         string
	GrpcProtocol string
	GrpcPort     string
}

type config struct {
	serverConfig *ServerConfig
}

type Config interface {
	GetConfig() *config
	GetServerConfig() *ServerConfig
}
