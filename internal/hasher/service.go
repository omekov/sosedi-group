package hasher

type Service struct {
	Repository Hasher
}

func NewService() *Service {
	return &Service{
		Repository: NewHasherRepository(),
	}
}

type HasherService interface {
}
