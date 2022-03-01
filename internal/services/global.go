package services

import (
	"context"
	"log"
	"reflect"
)

type GlobalServiceKey string

var (
	globanContext context.Context
	// if you want to add more global services, please create key here
	// ExampleServiceKey = GlobalServiceKey("exampleService")
	AuthServiceKey       = GlobalServiceKey("authService")
	GrpcRoutesServiceKey = GlobalServiceKey("grpcRoutesService")

	// and then add the service in the following map to register it
	globalServicesMap = map[reflect.Type]GlobalServiceKey{
		reflect.TypeOf(&AuthService{}):       AuthServiceKey,
		reflect.TypeOf(&GrpcMethodService{}): GrpcRoutesServiceKey,
	}
)

func RegisterGlobalService(ctx context.Context, service interface{}) context.Context {
	serviceType := reflect.TypeOf(service)
	if _, ok := globalServicesMap[serviceType]; !ok {
		log.Fatalf("unknown global service: %v, please register following the instrucctions in global.go", serviceType)
	}

	ctx = context.WithValue(ctx, globalServicesMap[serviceType], service)
	globanContext = ctx
	return ctx
}

func GetGlobalService(k GlobalServiceKey) interface{} {
	v := globanContext.Value(k)
	if v == nil {
		log.Fatalf("not found value: %v", k)
	}
	return v
}
