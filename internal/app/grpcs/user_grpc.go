package grpcs

import (
	"context"
	"fmt"
	"log"
	"login_grpc/internal/models"
	"login_grpc/internal/services"
	"login_grpc/pkg"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/wire"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	AuthorizationHeader = "Authorization"
)

type UserServiceServer struct {
	pkg.UnimplementedUserServiceServer
	opts *UserServerOpts
}

type UserServerOpts struct {
	UserService *services.UserService
}

func NewUserServiceServer(opts *UserServerOpts) *UserServiceServer {
	return &UserServiceServer{
		opts: opts,
	}
}

var UserServiceServerSet = wire.NewSet(wire.Struct(new(UserServerOpts), "*"), NewUserServiceServer)

func (s *UserServiceServer) GetUserInfo(ctx context.Context, in *pkg.GetUserInfoRequest) (*pkg.GetUserInfoResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Fatalf("Failed to get metadata")
	}

	authorization := md.Get(AuthorizationHeader)

	fmt.Printf("Authorization: %v\n", authorization)

	token, err := jwt.Parse(authorization[0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return "aasd", nil
	})

	if !token.Valid {
		return nil, status.Error(codes.PermissionDenied, "Token is not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["Username"], claims["Id"])
	} else {
		fmt.Println(err)
	}

	user, _ := s.opts.UserService.GetUserByUsername(in.Username)

	if user.ID == 0 {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	return &pkg.GetUserInfoResponse{
		Username: user.Username,
		Id:       int32(user.ID),
		CreateAt: user.CreatedAt.Local().String(),
	}, nil
}

func (s *UserServiceServer) CreateUser(ctx context.Context, in *pkg.CreateUserRequest) (*pkg.CreateUserResponse, error) {
	user := models.User{
		Username:  in.Username,
		Password:  in.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := s.opts.UserService.CreateUser(user)

	fmt.Println(result)

	// TODO: Handle error
	if err != nil {
		return &pkg.CreateUserResponse{
			Success: false,
		}, err
	}

	return &pkg.CreateUserResponse{
		Success: true,
	}, nil
}
