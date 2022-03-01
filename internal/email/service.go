package firstname

import "context"

type Service struct {
	Repository FirstName
}

func NewService() *Service {
	return &Service{
		Repository: NewSubstrRepository(),
	}
}

type FirstNameService interface {
	FindBySymbol(ctx context.Context, query string) (string, error)
	FindBySelf(ctx context.Context, query string) ([]string, error)
}
