package services

import (
	"context"
	"log"
	"reflect"
)

type GlobalServiceKey string

var (
	GlobalContext context.Context
	// if you want to add more global services, please create key here
	// ExampleServiceKey = GlobalServiceKey("exampleService")
	AuthServiceKey       = GlobalServiceKey("authService")
	GrpcRoutesServiceKey = GlobalServiceKey("grpcRoutesService")

	// and then add the service in the following map to register it
	GlobalServicesMap = map[reflect.Type]GlobalServiceKey{
		reflect.TypeOf(&AuthService{}):       AuthServiceKey,
		reflect.TypeOf(&GrpcRoutesService{}): GrpcRoutesServiceKey,
	}
)

func RegisterGlobalService(ctx context.Context, service interface{}) context.Context {
	serviceType := reflect.TypeOf(service)
	if _, ok := GlobalServicesMap[serviceType]; !ok {
		log.Fatalf("unknown global service: %v, please register following the instrucctions in global.go", serviceType)
	}

	ctx = context.WithValue(ctx, GlobalServicesMap[serviceType], service)
	GlobalContext = ctx
	return ctx
}

func GetGlobalService(k GlobalServiceKey) interface{} {
	v := GlobalContext.Value(k)
	if v == nil {
		log.Fatalf("not found value: %v", k)
	}
	return v
}
