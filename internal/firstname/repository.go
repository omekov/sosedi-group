package firstname

import (
	"context"
)

type firstNameRepository struct {
}

func NewSubstrRepository() *firstNameRepository {
	return &firstNameRepository{}
}

type FirstName interface {
	GetAll(ctx context.Context) ([]string, error)
}

func (r *firstNameRepository) GetAll(ctx context.Context) ([]string, error) {
	return []string{}, nil
}
