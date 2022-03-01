package counter

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type couterRepository struct {
	client *redis.Client
	key    string
}

func NewCounterRepository(client *redis.Client) *couterRepository {
	return &couterRepository{client: client, key: "counter"}
}

type Counter interface {
	Set(ctx context.Context, num uint64) error
	Get(ctx context.Context) (string, error)
}

func (r *couterRepository) Set(ctx context.Context, num uint64) error {
	if r.key == "" {
		return fmt.Errorf("redisSet key: %s empty", r.key)
	}

	return r.client.Set(ctx, r.key, num, 0).Err()
}

func (r *couterRepository) Get(ctx context.Context) (string, error) {
	if r.key == "" {
		return "", fmt.Errorf("redisGet key: %s empty", r.key)
	}

	val, err := r.client.Get(ctx, r.key).Result()
	if err == redis.Nil {
		return val, fmt.Errorf("redisGet key: %s not exist", r.key)
	}

	return val, nil
}
