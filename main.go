package main

import (
	"log"
	postgresDb "login_grpc/pkg/database"
	"login_grpc/pkg/protogen/loginpb"
	loginGrpc "login_grpc/pkg/rpc/login"
	userGrpc "login_grpc/pkg/rpc/user"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	d := postgresDb.Init()

	s := grpc.NewServer()
	loginService := loginGrpc.NewLoginService(loginGrpc.WithDbClient(d))
	userService := userGrpc.NewUserService(userGrpc.WithDbClient(d))

	loginpb.RegisterLoginServiceServer(s, loginService)
	loginpb.RegisterUserServiceServer(s, userService)

	reflection.Register(s)

	log.Println("Starting server on port 50051")
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
