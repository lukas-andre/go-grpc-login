package common

import (
	"context"
	"log"
)

type GlobalService string

var GlobalContext context.Context

func SetGlobalService(ctx context.Context, k GlobalService, v interface{}) context.Context {
	ctx = context.WithValue(ctx, GlobalService(k), v)
	GlobalContext = ctx
	return ctx
}

func GetGlobalService(k GlobalService) interface{} {
	v := GlobalContext.Value(k)
	if v == nil {
		log.Fatalf("not found value: %v", k)
	}

	return v
}
