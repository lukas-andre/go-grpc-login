package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func LoggingServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// start := time.Now()

	h, err := handler(ctx, req)
	// grpclog.Infof("Request - Method:%s\tDuration:%s\tError:%v\n",
	// 	info.FullMethod,
	// 	time.Since(start),
	// 	err)
	return h, err
}
