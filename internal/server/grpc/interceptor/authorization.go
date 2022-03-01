package interceptor

import (
	"context"
	"errors"
	"login_grpc/internal/common"
	"login_grpc/internal/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

var (
	AuthorizationHeader = "Authorization"
)

func AuthorizationServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	tokenHandler := common.GetGlobalService(common.GlobalService("tokenHandler")).(*services.TokenHandler)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("failed to get metadata")
	}

	token := md.Get(AuthorizationHeader)
	grpclog.Infof("Some authorization:%s, some token:%s", info.FullMethod, token)

	_, err := tokenHandler.ParseToken(token[0])
	if err != nil {
		return nil, err
	}

	h, _ := handler(ctx, req)
	return h, err
}
