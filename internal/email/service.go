package email

type Service struct {
	Repository Email
}

func NewService() *Service {
	return &Service{
		Repository: NewEmailRepository(),
	}
}

type EmailService interface {
	FindEmail(text string) (string, error)
}

func (s *Service) FindEmail(text string) (string, error) {

	return "", nil
}
