package grpc

import (
	"fmt"
	"login_grpc/internal/config"
	"login_grpc/internal/server"
	"login_grpc/internal/server/grpc/interceptor"
	"login_grpc/pkg"
	"net"
	"os"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	config      *config.ServerConfig
	userServer  pkg.UserServiceServer
	loginServer pkg.LoginServiceServer
}

func NewServer(c *config.ServerConfig, us pkg.UserServiceServer, ls pkg.LoginServiceServer) server.Server {
	return &grpcServer{config: c, userServer: us, loginServer: ls}
}

func (s *grpcServer) Serve() error {
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.GrpcPort)
	listener, err := net.Listen(s.config.GrpcProtocol, addr)

	if err != nil {
		return err
	}

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	srv := grpc.NewServer(withUnaryInterceptor())

	pkg.RegisterUserServiceServer(srv, s.userServer)
	pkg.RegisterLoginServiceServer(srv, s.loginServer)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		return err
	}

	return nil
}

func withUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.LoggingServerInterceptor,
		interceptor.AuthorizationServerInterceptor,
	))
}
