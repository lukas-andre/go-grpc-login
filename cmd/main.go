package main

import (
	"context"
	"log"

	"login_grpc/internal/common"
	"login_grpc/internal/config"
	"login_grpc/internal/dao"
	di "login_grpc/internal/di_container"
	"login_grpc/internal/server/grpc"

	"golang.org/x/sync/errgroup"
)

func main() {
	config := config.ViperConfigInit()

	db, err := dao.NewPostgresDB(config.GetDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	var (
		userRepository = di.InitializeUserRepository(db)
		authRepository = di.InitializeAuthRepository(db)

		tokenHandler = di.InitializeTokenHandler(config.GetJWTConfig())

		userService = di.InitializeUserService(userRepository)
		authService = di.InitializeAuthService(authRepository, tokenHandler)

		userServiceServer = di.InitializeUserServiceServer(userService, authService)
		authServiceServer = di.InitializeAuthServiceServer(authService, userService, tokenHandler)
	)

	ctx := context.Background()
	ctx = common.SetGlobalService(ctx, common.GlobalService("tokenHandler"), tokenHandler)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		srv := grpc.NewServer(config.GetServerConfig(), userServiceServer, authServiceServer)

		log.Printf("gRPC server running at %s://%s:%s ...\n", "tcp", "0.0.0.0", "50051")
		return srv.Serve()
	})
	log.Fatal(g.Wait())

}
