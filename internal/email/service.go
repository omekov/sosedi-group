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
}
