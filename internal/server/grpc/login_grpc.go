package grpc

import (
	"context"
	"login_grpc/internal/services"
	"login_grpc/pkg"

	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type AuthServiceServer struct {
	pkg.UnimplementedLoginServiceServer
	opts AuthServerOptions
}

type AuthServerOptions struct {
	AuthService  *services.AuthService
	UserService  *services.UserService
	TokenHandler *services.TokenHandler
}

func NewAuthServiceServer(opts AuthServerOptions) *AuthServiceServer {
	return &AuthServiceServer{
		opts: opts,
	}
}

var AuthServiceServerSet = wire.NewSet(wire.Struct(new(AuthServerOptions), "*"), NewAuthServiceServer)

func (s *AuthServiceServer) Login(ctx context.Context, in *pkg.LoginRequest) (*pkg.LoginResponse, error) {
	user, err := s.opts.UserService.GetUserByUsername(in.Username)
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	token, err := s.opts.TokenHandler.CreateToken(user)
	if err != nil {
		grpclog.Errorf("Error creating token: %v", err)
		return nil, status.Error(codes.FailedPrecondition, "error creating token")
	}

	return &pkg.LoginResponse{
		Token: token,
	}, nil
}
