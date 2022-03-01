package interceptor

import (
	"context"
	"login_grpc/internal/services"

	"google.golang.org/grpc"
)

func AuthorizationServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	routesServices := services.GetGlobalService(services.GrpcRoutesServiceKey).(*services.GrpcRoutesService)

	if routesServices.IsPublicRoute(info.FullMethod) {
		return handler(ctx, req)
	}

	authService := services.GetGlobalService(services.AuthServiceKey).(*services.AuthService)
	_, err := authService.ValidateToken(ctx)

	if err != nil {
		return nil, err
	}

	h, _ := handler(ctx, req)
	return h, err
}
