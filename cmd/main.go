package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"login_grpc/internal/app/grpcs"
	"login_grpc/internal/app/grpcs/interceptor"
	"login_grpc/internal/dao"
	"login_grpc/internal/repository"
	"login_grpc/internal/services"
	"login_grpc/pkg"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func withUnaryInterceptor(f grpc.UnaryClientInterceptor) grpc.DialOption {
	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		interceptor.LoggingServerInterceptor,
		interceptor.AuthorizationServerInterceptor,
	))
}

func clientInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()
	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)
	// Logic after invoking the invoker
	log.Printf("Invoked RPC method=%s; Duration=%s; Error=%v", method,
		time.Since(start), err)
	return err
}

func main() {
	// DB
	d, err := dao.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// GRPC logs
	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	// GRPC server
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// Register repositories
	userRepository := repository.NewUserRepository(repository.WithUserRepositoryDao(d))

	// Register services
	userService := services.NewUserService(services.WithUserRepository(userRepository))
	// tokenHandler := services.NewTokenHandler("secret")

	// loginService := loginGrpc.New(loginGrpc.WithDbClient(d))
	userServerService := grpcs.NewUserServiceServer(grpcs.WithUserServerUserService(userService))

	// loginpb.RegisterLoginServiceServer(s, loginService)
	pkg.RegisterUserServiceServer(s, userServerService)

	reflection.Register(s)

	log.Println("Starting server on port 50051")
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
