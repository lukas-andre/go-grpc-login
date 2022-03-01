package services

type GrpcMethodService struct {
	publicMethods []*GrpcMethod
}

type GrpcMethod struct {
	Name string
}

func NewGrpMethodsService() *GrpcMethodService {
	return &GrpcMethodService{
		publicMethods: []*GrpcMethod{
			{
				Name: "/login.LoginService/Login",
			},
			{
				Name: "/login.UserService/CreateUser",
			},
		},
	}
}

func (s *GrpcMethodService) IsPublicMethod(method string) bool {
	isPublic := false
	for _, route := range s.publicMethods {
		if route.Name == method {
			isPublic = true
		}
	}
	return isPublic
}
