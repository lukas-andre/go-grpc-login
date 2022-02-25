package grpcs

import (
	"context"
	"login_grpc/internal/services"
	"login_grpc/pkg"

	"github.com/google/wire"
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
	return nil, nil
	// user := models.User{}

	// tx := s.dbClient.Find(&user, "username = ?", in.Username)

	// if tx.Error != nil {
	// 	return nil, nil
	// }

	// token, err := auth.CreateToken(user)

	// if err != nil {
	// 	return nil, err
	// }

	// return &pkg.LoginResponse{
	// 	Token: token,
	// }, nil
}
