package redis

import (
	"context"

	goredis "github.com/go-redis/redis/v8"
)

func NewClient(addr, pass string) *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr:     addr, // "localhost:6379"
		Password: pass, // no password set
		DB:       0,    // use default DB
	})
}

func Ping(ctx context.Context, client *goredis.Client) error {
	return client.Ping(ctx).Err()
}
