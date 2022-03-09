package email

import (
	"github.com/omekov/sosedi-group/pkg/util"
)

type Service struct {
	Repository Email
}

func NewService() *Service {
	return &Service{
		Repository: NewEmailRepository(),
	}
}

type EmailService interface {
	FindEmailFromText(text string) (string, error)
	FindIINFromText(text string) (string, error)
}

func (s *Service) FindEmailFromText(text string) (string, error) {
	email := util.FindEmailFromText(text)
	return email, nil
}

func (s *Service) FindIINFromText(text string) (string, error) {
	iin := util.FindIINFromText(text)
	return iin, nil
}
