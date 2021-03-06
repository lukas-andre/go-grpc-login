package interceptor

import (
	"context"
	"login_grpc/internal/services"

	"google.golang.org/grpc"
)

func AuthorizationServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	methodService := services.GetGlobalService(services.GrpcMethodsServiceKey).(*services.GrpcMethodService)
	if methodService.IsPublicMethod(info.FullMethod) {
		return handler(ctx, req)
	}

	authService := services.GetGlobalService(services.AuthServiceKey).(*services.AuthService)
	_, err := authService.ValidateToken(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
