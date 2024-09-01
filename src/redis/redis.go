package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.Background()

func Initialize(addr string) {
	// Initialize the redis client
	rdb = redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func GetClient() *redis.Client {
	return rdb
}

func GetContext() context.Context {
	return ctx
}
