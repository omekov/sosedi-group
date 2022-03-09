package hasher

import (
	"fmt"
	"hash/crc64"
	"time"
)

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

func (s *Service) Hasher(text string) {
	hash := crc64.New(crc64.MakeTable(crc64.ECMA))
	time.Now().UnixNano()
}