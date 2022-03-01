package main

import (
	"context"
	"log"
	"login_grpc/internal/config"
	"login_grpc/internal/dao"
	di "login_grpc/internal/di_container"
	"login_grpc/internal/server/grpc"
	"login_grpc/internal/services"

	"golang.org/x/sync/errgroup"
)

func main() {
	config := config.ViperConfigInit()

	db, err := dao.NewPostgresDB(config.GetDatabaseConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	var (
		routesServices = services.NewGrpMethodsService()

		userRepository = di.InitializeUserRepository(db)
		authRepository = di.InitializeAuthRepository(db)

		tokenHandler = di.InitializeTokenHandler(config.GetJWTConfig())

		userService = di.InitializeUserService(userRepository)
		authService = di.InitializeAuthService(authRepository, tokenHandler)

		userServiceServer = di.InitializeUserServiceServer(userService, authService)
		authServiceServer = di.InitializeAuthServiceServer(authService, userService, tokenHandler)
	)

	ctx := context.Background()

	ctx = services.RegisterGlobalService(ctx, authService)
	ctx = services.RegisterGlobalService(ctx, routesServices)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		sc := config.GetServerConfig()
		srv := grpc.NewServer(sc, userServiceServer, authServiceServer)
		log.Printf("gRPC server running at %s://%s:%s ...\n", sc.GrpcProtocol, sc.Host, sc.GrpcPort)
		return srv.Serve()
	})
	log.Fatal(g.Wait())

}
