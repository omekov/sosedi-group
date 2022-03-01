package counter

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/omekov/sosedi-group/pkg/util"
)

type Service struct {
	Repository Counter
}

func NewService(client *redis.Client) *Service {
	return &Service{
		Repository: NewCounterRepository(client),
	}
}

type CounterService interface {
	Get(ctx context.Context) (string, error)
	SetIncrement(ctx context.Context, num string) error
	SetDecrement(ctx context.Context, num string) error
}

func (s *Service) SetIncrement(ctx context.Context, num string) error {
	numUint, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		return err
	}

	numUint = util.Increment(numUint)
	return s.Repository.Set(ctx, numUint)
}

func (s *Service) SetDecrement(ctx context.Context, num string) error {
	numUint, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		return err
	}

	numUint = util.Decrement(numUint)

	return s.Repository.Set(ctx, numUint)
}

func (s *Service) Get(ctx context.Context) (string, error) {
	val, err := s.Repository.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("Get repository %s", err)
	}

	_, err = strconv.ParseUint(val, 10, 32)
	if err != nil {
		return "", fmt.Errorf("strconv.ParseUint %s", err)
	}

	return val, nil
}
