package grpcs

// import (
// 	"context"
// 	auth "login_grpc/pkg/auth"
// 	usersModels "login_grpc/pkg/models"
// 	"login_grpc/pkg/protogen/loginpb"

// 	"gorm.io/gorm"
// )

// type LoginServiceServer struct {
// 	loginpb.UnimplementedLoginServiceServer
// 	dependencies
// }

// type Dependencies func(*dependencies)

// type dependencies struct {
// 	dbClient *gorm.DB
// }

// func New(deps ...Dependencies) *LoginServiceServer {
// 	d := &dependencies{}
// 	for _, dep := range deps {
// 		dep(d)
// 	}
// 	return &LoginServiceServer{
// 		dependencies: *d,
// 	}
// }

// func WithDbClient(dbClient *gorm.DB) Dependencies {
// 	return func(d *dependencies) {
// 		d.dbClient = dbClient
// 	}
// }

// func (s *LoginServiceServer) Login(ctx context.Context, in *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
// 	user := usersModels.User{}

// 	tx := s.dbClient.Find(&user, "username = ?", in.Username)

// 	if tx.Error != nil {
// 		return nil, nil
// 	}

// 	token, err := auth.CreateToken(user)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &loginpb.LoginResponse{
// 		Token: token,
// 	}, nil
// }
