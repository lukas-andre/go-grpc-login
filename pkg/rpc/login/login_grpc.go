package loginGrpc

import (
	"context"
	"login_grpc/pkg/protogen/loginpb"

	usersModels "login_grpc/pkg/models"

	"gorm.io/gorm"
)

type LoginServiceServer struct {
	loginpb.UnimplementedLoginServiceServer
	dependencies
}

type Dependencies func(*dependencies)

type dependencies struct {
	dbClient *gorm.DB
}

func NewLoginService(deps ...Dependencies) *LoginServiceServer {
	d := &dependencies{}
	for _, dep := range deps {
		dep(d)
	}
	return &LoginServiceServer{
		dependencies: *d,
	}
}

func WithDbClient(dbClient *gorm.DB) Dependencies {
	return func(d *dependencies) {
		d.dbClient = dbClient
	}
}

func (s *LoginServiceServer) Login(ctx context.Context, in *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	user := usersModels.User{}

	s.dbClient.Find(&user, "username = ?", in.Username)

	return &loginpb.LoginResponse{
		Token: "TODO: CREATE A TOKEN",
	}, nil
}
