package services

type GrpcRoutesService struct {
	publicRoutes []*GrpcRoute
}

type GrpcRoute struct {
	Method string
}

func NewGrpcRoutesService() *GrpcRoutesService {
	return &GrpcRoutesService{
		publicRoutes: []*GrpcRoute{
			{
				Method: "/login.LoginService/Login",
			},
			{
				Method: "/login.UserService/CreateUser",
			},
		},
	}
}

func (s *GrpcRoutesService) IsPublicRoute(method string) bool {
	isPublic := false
	for _, route := range s.publicRoutes {
		if route.Method == method {
			isPublic = true
		}
	}
	return isPublic
}
