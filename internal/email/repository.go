package email

type emailRepository struct {
}

func NewEmailRepository() *emailRepository {
	return &emailRepository{}
}

type Email interface {
}
