package main

import (
	"log"
	"net"
	"os"

	"login_grpc/internal/config"
	"login_grpc/internal/dao"
	di "login_grpc/internal/di_container"
	"login_grpc/pkg"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

// func withUnaryInterceptor(f grpc.UnaryClientInterceptor) grpc.DialOption {
// 	return grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
// 		interceptor.LoggingServerInterceptor,
// 		interceptor.AuthorizationServerInterceptor,
// 	))
// }

// func clientInterceptor(
// 	ctx context.Context,
// 	method string,
// 	req interface{},
// 	reply interface{},
// 	cc *grpc.ClientConn,
// 	invoker grpc.UnaryInvoker,
// 	opts ...grpc.CallOption,
// ) error {
// 	// Logic before invoking the invoker
// 	start := time.Now()
// 	// Calls the invoker to execute RPC
// 	err := invoker(ctx, method, req, reply, cc, opts...)
// 	// Logic after invoking the invoker
// 	log.Printf("Invoked RPC method=%s; Duration=%s; Error=%v", method,
// 		time.Since(start), err)
// 	return err
// }

func main() {
	// TODO: CREATE CONFIG STRUCT AND RETURN IT FROM THIS FUNCTION
	// SO WE ARE NOT GOING TO DEPEND OF THE CONFIG READER
	c := config.Init()

	d, err := dao.NewPostgresDB(c)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	// Dependency Injection Container
	userRepository := di.InitializeUserRepository(d)
	authRepository := di.InitializeAuthRepository(d)

	userService := di.InitializeUserService(userRepository)
	authService := di.InitializeAuthService(authRepository)

	tokenHandler := di.InitializeTokenHandler(c.GetString("jwt.secret"))

	userServerService := di.InitializeUserServiceServer(userService)
	authServiceServer := di.InitializeAuthServiceServer(authService, userService, tokenHandler)

	s := grpc.NewServer()
	pkg.RegisterUserServiceServer(s, userServerService)
	pkg.RegisterLoginServiceServer(s, authServiceServer)

	reflection.Register(s)

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting server on port 50051")
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
