package main

import (
	"context"
	"github.com/nitishm/go-rejson/v4"
	"github.com/redis/go-redis/v9"
)

var (
	ctx context.Context
	rh  *rejson.Handler
	rbd *redis.Client
)

func init() {
	ctx := context.Background()
	rbd := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	rh = rejson.NewReJSONHandler()
	rh.SetGoRedisClientWithContext(ctx, rbd)
}
